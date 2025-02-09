package redis

import (
	"github.com/go-redis/redis/v7"
	"github.com/qingyggg/blog_server/pkg/constants"
	"time"
)

var (
	ExpireTime = time.Hour * 3
	//ExpireTime  = time.Second
	rdbCollect    *redis.Client
	rdbFavorite   *redis.Client
	rdbComment    *redis.Client
	rdbView       *redis.Client
	rdbUser       *redis.Client
	rdbUserOnline *redis.Client
)

func InitRedis() {
	rdbCollect = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       0,
	})
	rdbFavorite = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       1,
	})
	rdbComment = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       2,
	})
	rdbView = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       3,
	})
	//uid:uHashId,用以减小请求User数据库的次数
	rdbUser = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       4,
	})
	//uhashId:1
	rdbUserOnline = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       5,
	})
}
