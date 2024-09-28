package mongo

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	CommentCol    *mongo.Collection
	CmtClosureCol *mongo.Collection
)

func InitComment() {
	CommentCol = DB.Collection("comment") //存储文章评论的表
	CmtClosureCol = DB.Collection("comment_closure")

	CreateIndex(CommentCol, bson.M{"hash_id": 1}, true)
	CreateIndex(CommentCol, bson.M{"article_id": 1}, false)
	CreateIndex(CommentCol, bson.M{"article_id": 1, "hash_id": 1}, true)

	CreateIndex(CmtClosureCol, bson.M{"ancestor": 1}, false)
	CreateIndex(CmtClosureCol, bson.M{"descendant": 1}, false)
}

type Comment struct {
	ArticleID  string        `bson:"article_id"`
	Content    string        `bson:"content"`
	Degree     int8          `bson:"degree"` //用以请求顶级评论,degree==1
	UserID     string        `bson:"user_id"`
	HashID     string        `bson:"hash_id"`
	ParentID   string        `bson:"parent_id"`
	CreateTime bson.DateTime `bson:"create_time"`
}

type CommentClosure struct {
	AncestorID   string `bson:"ancestor"`
	DescendantID string `bson:"descendant"`
	Depth        int    `bson:"depth"`
}

// CommentItem 用以请求comment list
type CommentItem struct {
	ArticleID  string        `bson:"article_id"`
	Content    string        `bson:"content"`
	Degree     int8          `bson:"degree"`
	UserID     string        `bson:"user_id"`
	HashID     string        `bson:"hash_id"`
	CreateTime bson.DateTime `bson:"create_time"`
	ChildNum   int64         `bson:"child_num"`
}

type CommentItemForSub struct {
	ArticleID  string        `bson:"article_id"`
	Content    string        `bson:"content"`
	Degree     int8          `bson:"degree"`
	UserID     string        `bson:"user_id"`
	HashID     string        `bson:"hash_id"`
	ParentUID  string        `bson:"parent_uid"` //@xxx reply_to @xxx
	CreateTime bson.DateTime `bson:"create_time"`
}
