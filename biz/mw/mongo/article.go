package mongo

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var ArticleCollection *mongo.Collection

// Article 定义 article 数据结构
type Article struct {
	ArticleID string `bson:"article_id"`
	Content   string `bson:"content"`
}

func InitArticle() {
	ArticleCollection = DB.Collection("article") //存储文章内容的表
	CreateIndex(ArticleCollection, bson.M{"article_id": 1}, true)
}
