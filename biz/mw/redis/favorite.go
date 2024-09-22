package redis

import (
	"strconv"
)

const (
	likeSuffix     = ":like"
	likedSuffix    = ":liked"
	dislikeSuffix  = ":dislike"
	dislikedSuffix = ":disliked"
	msgLike        = "like"
	msgDislike     = "dislike"
)

type (
	Favorite struct{}
)

func (f Favorite) AddLike(user_id, article_id int64) {
	add(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix, article_id)
}

func (f Favorite) AddLiked(user_id, article_id int64) {
	add(rdbFavorite, strconv.FormatInt(article_id, 10)+likedSuffix, user_id)
}

func (f Favorite) AddDislike(user_id, article_id int64) {
	add(rdbFavorite, strconv.FormatInt(user_id, 10)+dislikeSuffix, article_id)
}

func (f Favorite) AddDisLiked(user_id, article_id int64) {
	add(rdbFavorite, strconv.FormatInt(article_id, 10)+dislikedSuffix, user_id)
}

func (f Favorite) DelLike(user_id, article_id int64) {
	del(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix, article_id)
}

func (f Favorite) DelLiked(user_id, article_id int64) {
	del(rdbFavorite, strconv.FormatInt(article_id, 10)+likedSuffix, user_id)
}
func (f Favorite) DelDisLike(user_id, article_id int64) {
	del(rdbFavorite, strconv.FormatInt(user_id, 10)+dislikeSuffix, article_id)
}

func (f Favorite) DelDisLiked(user_id, article_id int64) {
	del(rdbFavorite, strconv.FormatInt(article_id, 10)+dislikedSuffix, user_id)
}

func (f Favorite) CheckLike(user_id int64) bool {
	return check(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix)
}

func (f Favorite) CheckLiked(article_id int64) bool {
	return check(rdbFavorite, strconv.FormatInt(article_id, 10)+likedSuffix)
}

func (f Favorite) CheckDisLike(user_id int64) bool {
	return check(rdbFavorite, strconv.FormatInt(user_id, 10)+dislikeSuffix)
}

func (f Favorite) CheckDisLiked(article_id int64) bool {
	return check(rdbFavorite, strconv.FormatInt(article_id, 10)+dislikedSuffix)
}

func (f Favorite) ExistLike(user_id, article_id int64) bool {
	return exist(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix, article_id)
}

func (f Favorite) ExistLiked(user_id, article_id int64) bool {
	return exist(rdbFavorite, strconv.FormatInt(article_id, 10)+likedSuffix, user_id)
}

func (f Favorite) ExistDisLike(user_id, article_id int64) bool {
	return exist(rdbFavorite, strconv.FormatInt(user_id, 10)+dislikeSuffix, article_id)
}

func (f Favorite) ExistDisLiked(user_id, article_id int64) bool {
	return exist(rdbFavorite, strconv.FormatInt(article_id, 10)+dislikedSuffix, user_id)
}

func (f Favorite) CountLike(user_id int64) (int64, error) {
	return count(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix)
}

func (f Favorite) CountLiked(article_id int64) (int64, error) {
	return count(rdbFavorite, strconv.FormatInt(article_id, 10)+likedSuffix)
}

func (f Favorite) CountDisLike(user_id int64) (int64, error) {
	return count(rdbFavorite, strconv.FormatInt(user_id, 10)+dislikeSuffix)
}

func (f Favorite) CountDisLiked(article_id int64) (int64, error) {
	return count(rdbFavorite, strconv.FormatInt(article_id, 10)+dislikedSuffix)
}

func (f Favorite) GetLike(user_id int64) []int64 {
	return get(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix)
}

func (f Favorite) GetLiked(article_id int64) []int64 {
	return get(rdbFavorite, strconv.FormatInt(article_id, 10)+likedSuffix)
}

func (f Favorite) GetDisLike(user_id int64) []int64 {
	return get(rdbFavorite, strconv.FormatInt(user_id, 10)+dislikeSuffix)
}

func (f Favorite) GetDisLiked(article_id int64) []int64 {
	return get(rdbFavorite, strconv.FormatInt(article_id, 10)+dislikedSuffix)
}

func (f Favorite) CheckLikeStatus(user_id, article_id int64) string {
	if f.ExistLike(user_id, article_id) {
		return msgLike
	} else if f.ExistDisLike(user_id, article_id) {
		return msgDislike
	} else {
		return "no status"
	}
}

func (f Favorite) CheckLikedStatus(user_id, article_id int64) string {
	if f.ExistLiked(user_id, article_id) {
		return msgLike
	} else if f.ExistDisLiked(user_id, article_id) {
		return msgDislike
	} else {
		return "no status"
	}
}

func (f Favorite) FlushLikeStatus(user_id, article_id int64) {
	msg := f.CheckLikeStatus(user_id, article_id)
	if msg == msgLike {
		f.DelLike(user_id, article_id)
	} else if msg == msgDislike {
		f.DelDisLike(user_id, article_id)
	}
}

func (f Favorite) FlushLikedStatus(user_id, article_id int64) {
	msg := f.CheckLikedStatus(user_id, article_id)
	if msg == msgLike {
		f.DelLiked(user_id, article_id)
	} else if msg == msgDislike {
		f.DelDisLiked(user_id, article_id)
	}
}
