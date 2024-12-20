package db

import (
	"context"
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	"github.com/qingyggg/blog_server/biz/model/query"
	"github.com/qingyggg/blog_server/biz/mw/mongo"
	"github.com/qingyggg/blog_server/pkg/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateArticle(aInfo *orm_gen.Article, aContent string) error {
	err := QDB.Transaction(func(tx *query.Query) error {
		if err := tx.Article.Create(aInfo); err != nil {
			return err
		}
		_, err := mongo.ArticleCollection.InsertOne(context.TODO(), &mongo.Article{ArticleID: utils.ConvertByteHashToString(aInfo.HashID), Content: aContent})
		return err //nil or err
	})
	return err
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
		_, err := mongo.ArticleCollection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return utils.ConvertByteHashToString(aInfo.HashID), err
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
		_, err := mongo.ArticleCollection.DeleteOne(context.TODO(), filter)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

// GetArticleInfos: 获取一系列的文章卡片
func GetArticleInfos(uid string, offset int) (aInfos []*orm_gen.Article, err error) {
	var a = query.Article
	var expr query.IArticleDo
	if uid != "" {
		expr = a.Where(a.UserID.Eq(utils.ConvertStringHashToByte(uid)))
	} else {
		expr = a.Where()
	}
	aInfos, err = expr.Order(a.ID.Desc()).Limit(15).Offset(offset).Find()
	if err != nil {
		return nil, err
	}
	return aInfos, nil
}

// TakeArticle: 获取文章info以及内容
func TakeArticle(aHashID string) (aInfo *orm_gen.Article, content string, err error) {
	var a = query.Article
	aInfo, err = a.Where(a.HashID.Eq(utils.ConvertStringHashToByte(aHashID))).Take()
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
func GetWorkCount(uHashId string) (count int64, err error) {
	var a = query.Article
	count, err = a.Where(a.UserID.Eq(utils.ConvertStringHashToByte(uHashId))).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
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

// CheckArticleExistByHashId 根据文章的hash_id 字段查找文章
func CheckArticleExistByHashId(ahashId string) (bool, error) {
	var a = query.Article
	count, err := a.Where(a.HashID.Eq(utils.ConvertStringHashToByte(ahashId))).Count()
	if err != nil {
		return false, err
	} else if count != 0 {
		return true, nil
	} else {
		return false, nil
	}
}
