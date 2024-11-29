package db

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	"github.com/qingyggg/blog_server/biz/model/query"
	"github.com/qingyggg/blog_server/biz/mw/redis"
	"github.com/qingyggg/blog_server/pkg/utils"
	"gorm.io/gorm/clause"
	"strings"
	"time"
)

// redis:过期时间一小时+ttl+定时同步（10分钟一次）
var avrd redis.View

func ViewCountExist(ahashId string) (error, bool) {
	var err error
	var av = query.ArticleView
	//先检查redis，有就true
	err, exist := avrd.CheckViewCt(ahashId)
	if err != nil {
		return err, false
	}
	if exist {
		return nil, true
	}
	//木有再检查mysql，有就true,并且再同步进redis
	view, err := av.Where(av.ArticleID.Eq(utils.ConvertStringHashToByte(ahashId))).Take()
	if err != nil && err.Error() != "record not found" {
		return err, false
	}
	if view != nil {
		err := avrd.ViewCtAssign(ahashId, view.ViewCount) //redis同步
		if err != nil {
			return err, false
		}
		return nil, true
	}
	//mysql没有记录没有返回false
	return nil, false
}

func ViewCountGet(ahashId string) (error, int64) {
	return avrd.CountView(ahashId)
}

func ViewCountGets(ahashId []string) (error, map[string]int64) {
	var aidnoCacheByte [][]byte
	var av = query.ArticleView
	//先检查redis里的是否存在，存在的话再返回
	err, cMap := avrd.GetViewMap(ahashId)
	if err != nil {
		return err, nil
	}
	//将缓存没有命中的值收起来，去请求数据库的值
	for k, v := range cMap {
		if v == -1 {
			aidnoCacheByte = append(aidnoCacheByte, utils.ConvertStringHashToByte(k))
		}
	}
	//如果没有被缓存的值存在的话：
	if len(aidnoCacheByte) > 0 {
		avs, err := av.Where(av.ArticleID.In(aidnoCacheByte...)).Find()
		if err != nil {
			return err, nil
		}
		txPipe := avrd.GetViewClient().TxPipeline()
		for _, v := range avs {
			aid := utils.ConvertByteHashToString(v.ArticleID)
			cMap[aid] = v.ViewCount
			txPipe.Set(aid, v.ViewCount, redis.ExpireTime)
		}
		//请求完毕后，将值缓存进redis里，等待redis定时定时同步进数据库
		_, err = txPipe.Exec()
		if err != nil {
			return err, nil
		}
	}
	return nil, cMap
}

// ViewCountInit 在创建文章后调用该函数，在redis里初始化文章阅读数记录
func ViewCountInit(ahashId string) error {
	var va = query.ArticleView
	//这里初始化redis,下面再初始化数据库的
	err := avrd.ViewCtAssign(ahashId, 0)
	if err != nil {
		return err
	}
	err = va.Create(&orm_gen.ArticleView{ArticleID: utils.ConvertStringHashToByte(ahashId), ViewCount: 0})
	if err != nil {
		return err
	}
	return nil
}
func ViewCountIncr(ahashId string) error {
	//使用redis递加，并且等待redis定时同步进数据库
	err := avrd.IncrView(ahashId)
	return err
}

// ViewCountDel 删除文章的时候，调用该函数
func ViewCountDel(ahashId string) error {
	var err error
	var av = query.ArticleView
	err, exist := avrd.CheckViewCt(ahashId)
	if err != nil {
		return err
	}
	if exist {
		//清理redis
		if err := avrd.DelViewCt(ahashId); err != nil {
			return err
		}
	}
	//清理数据库
	_, err = av.Where(av.ArticleID.Eq(utils.ConvertStringHashToByte(ahashId))).Delete()
	return err
}

// StartPeriodicSyncForViewCt 定期任务，每 10 分钟将 Redis 中的view count数据同步到 MySQL
func StartPeriodicSyncForViewCt() {
	var vrdb redis.View
	var va = query.ArticleView
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		// 扫描 Redis 中所有与文章阅读数相关的 key
		var cursor uint64
		for {
			keys, cursor, err := vrdb.GetViewClient().Scan(cursor, fmt.Sprintf("*%s", redis.ViewCountSuffix), 100).Result()
			if err != nil {
				hlog.Fatal("error to scan redis keys--> for view count")
				break
			}

			//遍历每个 key，并同步数据到 MySQL
			var aHashIds []string
			for _, key := range keys {
				//从 key 中解析出 articleID
				aHashId := strings.Split(key, ":")[0]
				aHashIds = append(aHashIds, aHashId)
			}
			err, viewMap := vrdb.GetViewMap(aHashIds) //use redis pipeline
			if err != nil {
				hlog.Fatal("error to get redis viewMap")
			}
			var views []*orm_gen.ArticleView
			for k, v := range viewMap {
				views = append(views, &orm_gen.ArticleView{ArticleID: utils.ConvertStringHashToByte(k), ViewCount: v})
			}
			err = va.Clauses(clause.OnConflict{DoUpdates: clause.AssignmentColumns([]string{"view_count"})}).CreateInBatches(views, len(views))
			if err != nil {
				hlog.Fatal(err)
			}
			// 如果游标为 0，说明扫描完毕
			if cursor == 0 {
				break
			}
		}
	}
}
