package mongo

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	Client            *mongo.Client
	DB                *mongo.Database
	ArticleCollection *mongo.Collection
)

func Init() {
	var err error
	// 定义MongoDB的用户名和密码
	username := "blog_server"
	password := "blog_server123"
	host := "localhost"
	port := "18006"

	// 构造MongoDB连接URI，包含用户名和密码
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?connect=direct", username, password, host, port)
	clientOpts := options.Client().ApplyURI(
		uri)
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
	ArticleCollection = DB.Collection("article")
}

// Article 定义 article 数据结构
type Article struct {
	ArticleID string `bson:"article_id"`
	Content   string `bson:"content"`
}
