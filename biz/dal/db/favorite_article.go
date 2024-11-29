package db

import (
	"errors"
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	"github.com/qingyggg/blog_server/biz/model/query"
	"github.com/qingyggg/blog_server/biz/mw/redis"
	"github.com/qingyggg/blog_server/pkg/utils"
	"gorm.io/gorm"
	"sync"
)

// register redis operate strategy
var rdFavorite redis.Favorite

func AFavorite(aHashId string, uHashId string, status int32) error {
	var fa = query.ArticleFavorite
	err := fa.Create(&orm_gen.ArticleFavorite{
		ArticleID: utils.ConvertStringHashToByte(aHashId),
		UserID:    utils.ConvertStringHashToByte(uHashId),
		Status:    status,
	})
	if err != nil {
		return err
	}
	//redis同步
	errChan := make(chan error, 2)
	defer close(errChan)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if status == 1 {
			if err := rdFavorite.Like(uHashId, aHashId); err != nil {
				errChan <- err
				return
			}
		} else if status == 2 {
			if err := rdFavorite.Hate(uHashId, aHashId); err != nil {
				errChan <- err
				return
			}
		}
	}()
	go func() {
		defer wg.Done()
		if status == 1 {
			err, exist := rdFavorite.CheckLikeCt(aHashId)
			if err != nil {
				errChan <- err
				return
			}
			if exist {
				err = rdFavorite.IncrLike(aHashId)
				if err != nil {
					errChan <- err
					return
				}
			} else {
				err, count := aFavoriteCtGet(aHashId)
				if err != nil {
					errChan <- err
					return
				}
				err = rdFavorite.LikeCtAssign(aHashId, count)
				if err != nil {
					errChan <- err
					return
				}
			}
		}
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return err
		}
	default:
	}
	return nil
}

func AFavoriteExist(aHashId string, uHashId string) (err error, exSignal int32) {
	// signal: 1 -> like, 2 -> hate, 3 -> none

	// 获取 redis like
	err, has := rdFavorite.CheckLike(aHashId)
	if err != nil {
		return err, 3
	}
	if has {
		err, exist := rdFavorite.ExistLike(uHashId, aHashId)
		if err != nil {
			return err, 3
		}
		if exist {
			return nil, 1 // 用户点赞过
		}
	}

	// 获取 redis hate
	err, has = rdFavorite.CheckHate(aHashId)
	if err != nil {
		return err, 3
	}
	if has {
		err, exist := rdFavorite.ExistHate(uHashId, aHashId)
		if err != nil {
			return err, 3
		}
		if exist {
			return nil, 2 // 用户点踩过
		}
	}

	// 如果 redis 中没有记录，查询数据库并进行同步
	err, exSignal = aFavoriteExist(aHashId, uHashId)
	if err != nil {
		return err, 3
	}

	// 同步到 redis 中
	if exSignal == 1 {
		// 用户点赞
		if err = rdFavorite.Like(uHashId, aHashId); err != nil {
			return err, 3
		}
	} else if exSignal == 2 {
		// 用户点踩
		if err = rdFavorite.Hate(uHashId, aHashId); err != nil {
			return err, 3
		}
	}

	return nil, exSignal
}

func aFavoriteExist(aHashId string, uHashId string) (error, int32) {
	//找不到，从mysql里取
	var fa = query.ArticleFavorite
	af, err := fa.Where(fa.ArticleID.Eq(utils.ConvertStringHashToByte(aHashId)), fa.UserID.Eq(utils.ConvertStringHashToByte(uHashId))).Take()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 3 //没有记录
		}
		return err, 3 //数据库错误
	}
	return nil, af.Status //点赞或者踩
}
func AFavoriteStatusFlush(aHashId string, uHashId string, fluSignal int32) error { //signal:1:like 2:hate
	errChan := make(chan error, 2)
	defer close(errChan)
	var wg sync.WaitGroup
	var err error
	wg.Add(2)
	//数据库清除

	var fa = query.ArticleFavorite
	if _, err := fa.Where(fa.ArticleID.Eq(utils.ConvertStringHashToByte(aHashId)), fa.UserID.Eq(utils.ConvertStringHashToByte(uHashId))).Delete(); err != nil {
		return err
	}

	//redis清除
	go func() {
		if fluSignal == 1 {
			if err := rdFavorite.CancerLike(uHashId, aHashId); err != nil {
				errChan <- err
				return
			}
		} else if fluSignal == 2 {
			if err := rdFavorite.CancerHate(uHashId, aHashId); err != nil {
				errChan <- err
				return
			}
		}
		wg.Done()
	}()
	go func() {
		if fluSignal == 1 {
			err, exist := rdFavorite.CheckLikeCt(aHashId)
			if err != nil {
				errChan <- err
				return
			}
			if exist {
				err = rdFavorite.DecrLike(aHashId)
				if err != nil {
					errChan <- err
					return
				}
			} else {
				err, count := aFavoriteCtGet(aHashId)
				if err != nil {
					errChan <- err
					return
				}
				err = rdFavorite.LikeCtAssign(aHashId, count)
				if err != nil {
					errChan <- err
					return
				}
			}
		}

		wg.Done()
	}()
	select {
	case err = <-errChan:
		return err
	default:
	}
	return nil
}

// AFavoriteCtGet 获取文章的点赞数量
func AFavoriteCtGet(aHashId string) (error, int64) {
	//从redis获取count，redis没有mysql拿，并且做同步
	err, exist := rdFavorite.CheckLikeCt(aHashId)
	if err != nil {
		return err, 0
	}
	if exist {
		return rdFavorite.CountLike(aHashId)
	} else {
		err, count := aFavoriteCtGet(aHashId)
		if err != nil {
			return err, 0
		}
		//redis数据同步
		err = rdFavorite.LikeCtAssign(aHashId, count)
		if err != nil {
			return err, 0
		}
		return nil, count
	}

}

func aFavoriteCtGet(aHashId string) (error, int64) {
	var fa = query.ArticleFavorite
	count, err := fa.Where(fa.ArticleID.Eq(utils.ConvertStringHashToByte(aHashId)), fa.Status.Eq(1)).Count()
	if err != nil {
		return err, 0
	}
	return nil, count
}

func AFavoriteCtGetByAids(aHashIds []string) (error, map[string]int64) {
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
		err, exist := rdFavorite.CheckLikeCt(v)
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
			err, count := rdFavorite.CountLike(v)
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
			var fa = query.ArticleFavorite
			var results []struct {
				ArticleID []byte
				Count     int64
			}

			// 使用聚合查询来统计每个articleid的点赞数
			err := fa.Select(fa.ArticleID, fa.UserID.Count().As("Count")).
				Where(fa.ArticleID.In(aByteId...), fa.Status.Eq(1)).
				Group(fa.ArticleID).
				Scan(&results)
			if err != nil {
				errChan <- err
				return
			}

			// 更新结果到resMap，并将结果同步到Redis
			fpipe := rdFavorite.GetFavoriteClient().TxPipeline()
			for _, result := range results {
				aHashId := utils.ConvertByteHashToString(result.ArticleID) // 将byte转换回hash string
				mu.Lock()
				resMap[aHashId] = result.Count
				mu.Unlock()

				fpipe.Set(aHashId+redis.LikeCountSuffix, result.Count, redis.ExpireTime)
			}
			// 执行事务
			_, err = fpipe.Exec()
			if err != nil {
				errChan <- err
				return
			}
		}
	}()

	// 等待所有的goroutine完成任务
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

func AFavoriteDeleteByAid(aHashId string) error {
	var wg sync.WaitGroup
	var fa = query.ArticleFavorite
	wg.Add(3)
	errChan := make(chan error, 3)
	defer close(errChan)
	//数据库favorite
	go func() {
		defer wg.Done()
		_, err := fa.Where(fa.ArticleID.Eq(utils.ConvertStringHashToByte(aHashId))).Delete()
		if err != nil {
			errChan <- err
			return
		}
	}()
	//redis favorite
	go func() {
		defer wg.Done()
		err := rdFavorite.TruncateLikeStatus(aHashId) //favorite,like count
		if err != nil {
			errChan <- err
			return
		}
	}()
	//redis hate
	go func() {
		defer wg.Done()
		err := rdFavorite.TruncateHateStatus(aHashId)
		if err != nil {
			errChan <- err
			return
		}
	}()
	wg.Wait()
	select {
	case err := <-errChan:
		if err != nil {
			return err
		}
	default:
	}
	return nil
}
