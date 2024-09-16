package db

import (
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	"github.com/qingyggg/blog_server/biz/model/query"
)

var a = query.Article

func CreateArticle(article *orm_gen.Article) (aid int64, err error) {
	err = a.Create(article)
	if err != nil {
		return 0, err
	}
	return article.ID, err
}

//func GetArticlesByLastTime(lastTime time.Time) ([]*orm_gen.Article, error) {
//	videos := make([]*Video, constants.VideoFeedCount)
//	err := DB.Where("publish_time < ?", lastTime).Order("publish_time desc").Limit(constants.VideoFeedCount).Find(&videos).Error
//	if err != nil {
//		return videos, err
//	}
//	return videos, nil
//}

func GetArticlesByUserID(uid int64) (articles []*orm_gen.Article, err error) {
	articles, err = a.Where(a.UserID.Eq(uid)).Find()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

// GetWorkCount get the num of video published by the user
func GetWorkCount(uid int64) (count int64, err error) {
	count, err = a.Where(a.ID.Eq(uid)).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

// CheckArticleExistById  query if video exist
func CheckArticleExistById(aid int64) (bool, error) {
	count, err := a.Where(a.ID.Eq(aid)).Count()
	if err != nil {
		return false, err
	} else if count != 0 {
		return true, nil
	} else {
		return false, nil
	}
}
