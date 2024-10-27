package mongo

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/qingyggg/blog_server/pkg/constants"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	Client *mongo.Client
	DB     *mongo.Database
)

func Init() {
	var err error

	clientOpts := options.Client().ApplyURI(
		constants.MongoDefaultDSN)
	Client, err = mongo.Connect(clientOpts)
	//初始化数据库
	if err != nil {
		hlog.Fatal("mongoDB连接失败:", err)
	}
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		hlog.Fatal("mongoDB连接失败:", err)
	}

	hlog.Info("Connected to MongoDB!")

	// 选择数据库和集合（表）
	DB = Client.Database("blog_server")

	InitComment()
	InitArticle()
}
func CreateIndex(collection *mongo.Collection, field bson.M, needUnique bool) error {
	// 创建一个索引模型
	indexModel := mongo.IndexModel{
		Keys:    field,                                 // 升序索引
		Options: options.Index().SetUnique(needUnique), // 可选项，设置唯一索引
	}

	// 使用 Indexes().CreateOne 创建索引
	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}
