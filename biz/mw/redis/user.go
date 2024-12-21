package redis

import (
	"github.com/go-redis/redis/v7"
	"strconv"
)

type User struct{}

const (
	uidSuffix = ":uid"
)

func uget(c *redis.Client, k string) string {
	// 使用 GET 命令获取键对应的值
	val := c.Get(k).String()
	c.Expire(k, ExpireTime) //重置过期时间
	return val              // 返回获取到的值
}
func uset(c *redis.Client, k string, v string) error {
	tx := c.TxPipeline()
	tx.Set(k, v, ExpireTime)
	_, err := tx.Exec()
	return err
}
func UHashIdExist(uid int64) (error, bool) {
	return check(rdbUser, strconv.FormatInt(uid, 16))
}
func UHashIdGet(uid int64) string {
	return uget(rdbUser, strconv.FormatInt(uid, 16)+uidSuffix)
}
func UHashIdSet(uid int64, uHashId string) error {
	return uset(rdbUser, strconv.FormatInt(uid, 16)+uidSuffix, uHashId)
}
