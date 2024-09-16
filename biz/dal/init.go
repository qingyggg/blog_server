package dal

import (
	"github.com/qingyggg/blog_server/biz/dal/db"
	"github.com/qingyggg/blog_server/biz/model/query"
	"github.com/qingyggg/blog_server/biz/mw/redis"
)

// Init init dal
func Init() {
	db.Init() // mysql init
	query.SetDefault(db.DB)
	redis.InitRedis()
}
