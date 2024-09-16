package db

import (
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	"github.com/qingyggg/blog_server/biz/model/query"
	"github.com/qingyggg/blog_server/biz/mw/redis"
)

// register redis operate strategy
var rdFavorite redis.Favorite
var fa = query.ArticleFavorite

// AddNewFavorite add favorite relation
func AddNewFavorite(favorite *orm_gen.ArticleFavorite) (bool, error) {
	uid := favorite.UserID
	aid := favorite.ArticleID
	status := favorite.Status
	QDB.Transaction(func(tx *query.Query) error {
		err := tx.ArticleFavorite.Create(favorite)
		if err != nil {
			return err
		}
		if status == 1 {
			_, err := tx.Article.Where(tx.Article.ID.Eq(favorite.ArticleID)).UpdateSimple(a.LikeCount.Add(1))
			if err != nil {
				return err
			}
		}
		return nil
	})
	// add data to redis
	go func() {
		rdFavorite.FlushLikeStatus(uid, aid)
		rdFavorite.FlushLikedStatus(uid, aid)
		if favorite.Status == -1 { //-1踩 1点赞
			rdFavorite.AddDislike(uid, aid)
			rdFavorite.AddDisLiked(uid, aid)
		} else {
			rdFavorite.AddLike(uid, aid)
			rdFavorite.AddLiked(uid, aid)
		}
	}()
	return true, nil
}

// DeleteFavorite delete favorite relation
func DeleteFavorite(favorite *orm_gen.ArticleFavorite) (bool, error) {

	QDB.Transaction(func(tx *query.Query) error {
		_, err := tx.ArticleFavorite.Where(tx.ArticleFavorite.UserID.Eq(favorite.UserID), tx.ArticleFavorite.ArticleID.Eq(favorite.ArticleID)).Delete()
		if err != nil {
			return err
		}
		if favorite.Status == 1 {
			_, err := tx.Article.Where(tx.Article.ID.Eq(favorite.ArticleID)).UpdateSimple(a.LikeCount.Sub(1))
			if err != nil {
				return err
			}
		}
		return nil
	})
	go func() {
		rdFavorite.FlushLikeStatus(favorite.UserID, favorite.ArticleID)
		rdFavorite.FlushLikedStatus(favorite.UserID, favorite.ArticleID)
	}()
	return true, nil
}

// QueryFavoriteStatus query the like record by video_id and user_id
func QueryFavoriteStatus(uid, aid int64) (string, error) {
	msg := rdFavorite.CheckLikeStatus(uid, aid)
	if msg != "no status" {
		return msg, nil
	}
	msg = rdFavorite.CheckLikedStatus(uid, aid)
	if msg != "no status" {
		return msg, nil
	}
	fav, err := fa.Where(fa.UserID.Eq(uid), fa.ArticleID.Eq(aid)).Take()
	if err != nil {
		return "", err
	}
	if fav.Status == -1 {
		//踩
		return "dislike", nil
	} else {
		return "like", nil
	}
}

// QueryTotalFavoritedByAuthorID query the like num of all the video published by  the user
func QueryTotalFavoritedByAuthorID(auid int64) (int64, error) {
	var sum int64
	err := a.Select(a.LikeCount.Sum()).Where(a.UserID.Eq(auid)).Scan(&sum)
	if err != nil {
		return 0, err
	}
	return sum, nil
}

// getFavoriterIdList get the id list of liker of video in db
func getFavoriterIdList(aid int64) ([]int64, error) {
	favorite_actions, err := fa.Where(fa.ArticleID.Eq(aid), fa.Status.Eq(1)).Find()
	if err != nil {
		return nil, err
	}
	var result []int64
	for _, v := range favorite_actions {
		result = append(result, v.UserID)
	}
	return result, nil
}

// GetFavoriterIdList get the id list of liker of  video
func GetFavoriterIdList(aid int64) ([]int64, error) {
	if rdFavorite.CheckLiked(aid) {
		return rdFavorite.GetLiked(aid), nil
	}
	return getFavoriterIdList(aid)
}

// GetFavoriteCount count the favorite of video
func GetFavoriteCount(aid int64) (int64, error) {
	if rdFavorite.CheckLiked(aid) {
		return rdFavorite.CountLiked(aid)
	}
	// Not in the cache, go to the database to find and update the cache
	likes, err := getFavoriterIdList(aid)
	if err != nil {
		return 0, err
	}

	// update redis asynchronously
	go func(uids []int64, aid int64) {
		for _, uid := range uids {
			rdFavorite.AddLiked(uid, aid)
		}
	}(likes, aid)
	return int64(len(likes)), nil
}
