package mongo

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Client *mongo.Client

func Init() {
	var err error
	clientOpts := options.Client().ApplyURI(
		"mongodb://localhost:27017/?connect=direct")
	Client, err = mongo.Connect(clientOpts)
	if err != nil {
		hlog.Fatal("mongoDB连接失败:", err)
	} else {
		hlog.Info("成功连接mysql")
	}
}
