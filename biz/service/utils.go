package service_utils

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/qingyggg/blog_server/biz/dal/db"
	"github.com/qingyggg/blog_server/biz/mw/redis"
	"github.com/qingyggg/blog_server/pkg/utils"
)

// GetUHashId get uid from jwt
func GetUHashId(c *app.RequestContext) string {
	uid, exist := c.Get("current_user_id")
	if exist {
		//1.查找redis，找不到查找数据库，并且缓存进redis中
		err, ex := redis.UHashIdExist(uid.(int64))
		if err != nil {
			hlog.Warn("获取uhashId发生错误", err)
			return ""
		}
		if ex {
			return redis.UHashIdGet(uid.(int64))
		} else {
			user, err := db.QueryUserById(uid.(int64))
			if err != nil {
				hlog.Warn("获取uhashId发生错误", err)
				return ""
			}
			uHashId := utils.ConvertByteHashToString(user.HashID)
			go func() {
				err := redis.UHashIdSet(uid.(int64), uHashId)
				if err != nil {
					hlog.Warn("redis设置uhashId发生错误", err)
				}
			}()
			return uHashId
		}
	}
	return ""
}
