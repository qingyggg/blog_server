package redis

import (
	"github.com/go-redis/redis/v7"
	"github.com/qingyggg/blog_server/pkg/constants"
	"time"
)

var (
	ExpireTime = time.Hour * 3
	//ExpireTime  = time.Second
	rdbCollect  *redis.Client
	rdbFavorite *redis.Client
	rdbComment  *redis.Client
	rdbView     *redis.Client
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
}
