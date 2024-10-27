package db

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/qingyggg/blog_server/biz/model/query"
	"github.com/qingyggg/blog_server/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB
var QDB *query.Query

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		hlog.Fatal(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		hlog.Fatal(err)
	}
	QDB = query.Use(DB)
	hlog.Info("mysql数据库初始化成功")
}
