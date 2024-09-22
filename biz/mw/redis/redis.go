package redis

import (
	"github.com/go-redis/redis/v7"
	"github.com/qingyggg/blog_server/pkg/constants"
	"time"
)

var (
	expireTime  = time.Hour * 1
	rdbFollows  *redis.Client
	rdbFavorite *redis.Client
)

func InitRedis() {
	rdbFollows = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       0,
	})
	rdbFavorite = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       1,
	})
}
