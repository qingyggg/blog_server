package db

import (
	"context"
	"fmt"
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	"github.com/qingyggg/blog_server/biz/model/query"
	"github.com/qingyggg/blog_server/biz/mw/mongo"
	"github.com/qingyggg/blog_server/pkg/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"strconv"
	"time"
)

func CreateArticle(aInfo *orm_gen.Article, aContent string) (hashAid string, err error) {
	// 获取当前时间戳
	timestamp := time.Now().Unix()
	aHashID := utils.GetSHA256String(strconv.FormatInt(aInfo.UserID, 10) + strconv.FormatInt(timestamp, 10))
	aInfo.HashID = aHashID
	err = QDB.Transaction(func(tx *query.Query) error {
		if err := tx.Article.Create(aInfo); err != nil {
			return err
		}
		insertResult, err := mongo.ArticleCollection.InsertOne(context.TODO(), &mongo.Article{ArticleID: aHashID, Content: aContent})
		fmt.Println(insertResult)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return hashAid, err
}

func ModifyArticle(aInfo *orm_gen.Article, aContent string) (hashAid string, err error) {
	err = QDB.Transaction(func(tx *query.Query) error {
		_, err = tx.Article.Where(tx.Article.UserID.Eq(aInfo.UserID), tx.Article.HashID.Eq(aInfo.HashID)).Updates(aInfo)
		if err != nil {
			return err
		}
		filter := bson.M{"article_id": aInfo.HashID} // 过滤条件
		// 定义更新内容
		update := bson.M{
			"$set": bson.M{
				"content": aContent, // 更新 `content` 字段
			},
		}
		// 更新单个文档
		updateResult, err := mongo.ArticleCollection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return err
		}
		fmt.Println(updateResult)
		return nil
	})
	if err != nil {
		return "", err
	}

	return aInfo.HashID, err
}
func DeleteArticle(aInfo *orm_gen.Article) (err error) {
	err = QDB.Transaction(func(tx *query.Query) error {
		// 在 MySQL 中删除 article
		_, err = tx.Article.Where(tx.Article.UserID.Eq(aInfo.UserID), tx.Article.HashID.Eq(aInfo.HashID)).Delete()
		if err != nil {
			return err
		}

		// 在 MongoDB 中删除对应的 article
		filter := bson.M{"article_id": aInfo.HashID} // MongoDB 过滤条件，使用 HashID 匹配
		deleteResult, err := mongo.ArticleCollection.DeleteOne(context.TODO(), filter)
		if err != nil {
			return err
		}

		// 输出删除结果，供调试使用
		fmt.Println("MongoDB delete result:", deleteResult)

		return nil
	})

	return err
}

// GetArticleInfos: 获取一系列的文章卡片
func GetArticleInfos(uid int64, offset int) (aInfos []*orm_gen.Article, err error) {
	var a = query.Article
	var expr query.IArticleDo
	if uid != 0 {
		expr = a.Where(a.UserID.Eq(uid))
	} else {
		expr = a.Where()
	}
	aInfos, err = expr.Limit(10).Offset(offset).Find()
	if err != nil {
		return nil, err
	}
	return aInfos, nil
}

// TakeArticle: 获取文章info以及内容
func TakeArticle(aHashID string, uid int64) (aInfo *orm_gen.Article, content string, err error) {
	var a = query.Article
	aInfo, err = a.Where(a.UserID.Eq(uid), a.HashID.Eq(aHashID)).Take()
	if err != nil {
		return nil, "", err
	}
	filter := bson.M{"article_id": aHashID} // 过滤条件

	// 查找文章
	article := new(mongo.Article)
	err = mongo.ArticleCollection.FindOne(context.TODO(), filter).Decode(&article)
	if err != nil {
		return nil, "", err
	}
	content = article.Content
	return aInfo, content, nil
}

// GetWorkCount get the num of video published by the user
func GetWorkCount(uid int64) (count int64, err error) {
	var a = query.Article
	count, err = a.Where(a.ID.Eq(uid)).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CheckIsFavoriteByUid(uid int64, ahashId string) (bool, error) {
	//var a = query.Article
	return false, nil
}

func AddViewCount(uid int64, ahashId string) error {
	var a = query.Article
	_, err := a.Where(a.UserID.Eq(uid), a.HashID.Eq(ahashId)).UpdateSimple(a.ViewCount.Add(1))
	return err
}

// CheckArticleExistById  query if video exist
func CheckArticleExistById(aid int64) (bool, error) {
	var a = query.Article
	count, err := a.Where(a.ID.Eq(aid)).Count()
	if err != nil {
		return false, err
	} else if count != 0 {
		return true, nil
	} else {
		return false, nil
	}
}
