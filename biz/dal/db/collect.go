package db

import (
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	"github.com/qingyggg/blog_server/biz/model/query"
	"github.com/qingyggg/blog_server/biz/mw/redis"
	"github.com/qingyggg/blog_server/pkg/utils"
	"sync"
)

var rdCollect redis.Collect

func ACollectAdd(aHashId string, uHashId string, tag string) error {
	ac := query.ArticleCollect
	err := ac.Create(&orm_gen.ArticleCollect{ArticleID: utils.ConvertStringHashToByte(aHashId), UserID: utils.ConvertStringHashToByte(uHashId), Tag: tag})
	if err != nil {
		return err
	}
	err, exist := rdCollect.CheckCollectCt(aHashId)
	if err != nil {
		return err
	}
	if exist {
		err := rdCollect.IncrCollect(aHashId)
		if err != nil {
			return err
		}
	}
	return err
}

func ACollectDel(aHashId string, uHashId string) error {
	ac := query.ArticleCollect
	_, err := ac.Where(ac.ArticleID.Eq(utils.ConvertStringHashToByte(aHashId)), ac.UserID.Eq(utils.ConvertStringHashToByte(uHashId))).Delete()
	if err != nil {
		return err
	}
	err, exist := rdCollect.CheckCollectCt(aHashId)
	if err != nil {
		return err
	}
	if exist {
		err := rdCollect.DecrCollect(aHashId)
		if err != nil {
			return err
		}
	}
	return err
}

func ACollectDelByAid(aHashId string) error {
	//1.删除collect
	ac := query.ArticleCollect
	_, err := ac.Where(ac.ArticleID.Eq(utils.ConvertStringHashToByte(aHashId))).Delete()
	//2.删除collect count
	err, exist := rdCollect.CheckCollectCt(aHashId)
	if err != nil {
		return err
	}
	if exist {
		err := rdCollect.DelCollectCt(aHashId)
		if err != nil {
			return err
		}
	}
	return nil
}

func ACollectExist(aHashId string, uHashId string) (error, bool) {
	ac := query.ArticleCollect
	ct, err := ac.Where(ac.ArticleID.Eq(utils.ConvertStringHashToByte(aHashId)), ac.UserID.Eq(utils.ConvertStringHashToByte(uHashId))).Count()
	if err != nil {
		return err, false
	}
	if ct == 0 {
		return nil, false
	}
	return nil, true
}

func ACollectCtGet(aHashId string) (error, int64) {
	//从redis获取count，redis没有mysql拿，并且做同步
	err, exist := rdCollect.CheckCollectCt(aHashId)
	if err != nil {
		return err, 0
	}
	if exist {
		return rdCollect.CountCollect(aHashId)
	} else {
		err, count := aCollectCtGet(aHashId)
		if err != nil {
			return err, 0
		}
		//redis数据同步
		err = rdCollect.CollectCtAssign(aHashId, count)
		if err != nil {
			return err, 0
		}
		return nil, count
	}

}

func aCollectCtGet(aHashId string) (error, int64) {
	var ca = query.ArticleCollect
	count, err := ca.Where(ca.ArticleID.Eq(utils.ConvertStringHashToByte(aHashId))).Count()
	if err != nil {
		return err, 0
	}
	return nil, count
}

func ACollectCtGetByAids(aHashIds []string) (error, map[string]int64) {
	// 将所有的aid从Redis里获取，获取不到的aid收集起来，一起从数据库中获取，获取后将这些键值对存储进Redis
	var aidInCache []string // 存在于缓存里的aHashId:count
	var aidInNoCache []string
	var resMap = make(map[string]int64) // 初始化结果map
	var wg sync.WaitGroup
	var mu sync.Mutex              // 用于保护resMap的并发写入
	errChan := make(chan error, 2) // 处理错误的channel
	defer close(errChan)
	// 遍历输入的aHashIds，检查每个ID是否存在于缓存中
	for _, v := range aHashIds {
		err, exist := rdCollect.CheckCollectCt(v)
		if err != nil {
			return err, nil
		}
		if exist {
			aidInCache = append(aidInCache, v)
		} else {
			aidInNoCache = append(aidInNoCache, v)
		}
	}

	wg.Add(2) // 启动两个goroutine

	// 从缓存中获取点赞数的goroutine
	go func() {
		defer wg.Done()
		for _, v := range aidInCache {
			err, count := rdCollect.CountCollect(v)
			if err != nil {
				errChan <- err
				return
			}
			mu.Lock()
			resMap[v] = count
			mu.Unlock()
		}
	}()

	// 查询MySQL数据库并将结果同步到Redis的goroutine
	go func() {
		defer wg.Done()
		if len(aidInNoCache) > 0 {
			// 将HashID转换为数据库中的byte ID
			var aByteId [][]byte
			for _, v := range aidInNoCache {
				aByteId = append(aByteId, utils.ConvertStringHashToByte(v))
			}

			// 查询MySQL
			var ca = query.ArticleCollect
			var results []struct {
				ArticleID []byte
				Count     int64
			}

			// 使用聚合查询来统计每个articleid的点赞数
			err := ca.Select(ca.ArticleID, ca.UserID.Count().As("Count")).
				Where(ca.ArticleID.In(aByteId...)).
				Group(ca.ArticleID).
				Scan(&results)
			if err != nil {
				errChan <- err
				return
			}

			// 更新结果到resMap，并将结果同步到Redis
			fpipe := rdCollect.GetCollectClient().TxPipeline()
			for _, result := range results {
				aHashId := utils.ConvertByteHashToString(result.ArticleID) // 将byte转换回hash string
				mu.Lock()
				resMap[aHashId] = result.Count
				mu.Unlock()

				fpipe.Set(aHashId+redis.CollectCountSuffix, result.Count, redis.ExpireTime)
			}
			// 执行事务
			_, err = fpipe.Exec()
			if err != nil {
				errChan <- err
				return
			}
		}
	}()
	wg.Wait()
	// 检查是否有错误发生

	select {
	case err := <-errChan:
		if err != nil {
			return err, nil
		}
	default:

	}
	for _, v := range aHashIds {
		if _, exist := resMap[v]; exist == false {
			resMap[v] = 0
		}
	}
	return nil, resMap
}
