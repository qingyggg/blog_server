package redis

import "github.com/go-redis/redis/v7"

const (
	likeSuffix      = ":liked"
	hateSuffix      = ":hated"
	LikeCountSuffix = ":liked_count"
	HateCountSuffix = ":hated_count"
)

type Favorite struct{}

func (f Favorite) GetFavoriteClient() *redis.Client {
	return rdbFavorite
}

func (f Favorite) Like(uHashId, aHashId string) error {
	return sadd(rdbFavorite, aHashId+likeSuffix, uHashId)
}

func (f Favorite) CancerLike(uHashId, aHashId string) error {
	return srem(rdbFavorite, aHashId+likeSuffix, uHashId)
}

// CheckLike 该用户对文章是否点赞
func (f Favorite) ExistLike(uHashId, aHashId string) (error, bool) {
	return sexist(rdbFavorite, aHashId+likeSuffix, uHashId)
}

func (f Favorite) CheckLike(aHashId string) (error, bool) {
	return check(rdbFavorite, aHashId+likeSuffix)
}

func (f Favorite) Hate(uHashId, aHashId string) error {
	return sadd(rdbFavorite, aHashId+hateSuffix, uHashId)
}

func (f Favorite) CancerHate(uHashId, aHashId string) error {
	return srem(rdbFavorite, aHashId+hateSuffix, uHashId)
}

func (f Favorite) ExistHate(uHashId, aHashId string) (error, bool) {
	return sexist(rdbFavorite, aHashId+hateSuffix, uHashId)
}

func (f Favorite) CheckHate(aHashId string) (error, bool) {
	return check(rdbFavorite, aHashId+hateSuffix)
}

// 文章点赞处理逻辑
func (f Favorite) IncrLike(aHashId string) error {
	return incr(rdbFavorite, aHashId+LikeCountSuffix)
}

func (f Favorite) IncrHate(aHashId string) error {
	return incr(rdbFavorite, aHashId+HateCountSuffix)
}

func (f Favorite) DecrLike(aHashId string) error {
	return decr(rdbFavorite, aHashId+LikeCountSuffix)
}

func (f Favorite) DecrHate(aHashId string) error {
	return decr(rdbFavorite, aHashId+HateCountSuffix)
}

func (f Favorite) CountLike(aHashId string) (error, int64) {
	return nget(rdbFavorite, aHashId+LikeCountSuffix)
}

func (f Favorite) CountHate(aHashId string) (error, int64) {
	return nget(rdbFavorite, aHashId+HateCountSuffix)
}

//从数据库count点赞和踩的数量，并且赋值给redis

func (f Favorite) LikeCtAssign(aHashId string, count int64) error {
	return nset(rdbFavorite, aHashId+LikeCountSuffix, count)
}

func (f Favorite) HateCtAssign(aHashId string, count int64) error {
	return nset(rdbFavorite, aHashId+HateCountSuffix, count)
}

func (f Favorite) CheckLikeCt(aHashId string) (error, bool) {
	return check(rdbFavorite, aHashId+LikeCountSuffix)
}

func (f Favorite) CheckHateCt(aHashId string) (error, bool) {
	return check(rdbFavorite, aHashId+HateCountSuffix)
}

// 用户删除文章后，文章点赞依赖删除逻辑：
func (f Favorite) TruncateLikeStatus(aHashId string) error {
	err := delKey(rdbFavorite, aHashId+likeSuffix)
	if err != nil {
		return err
	}
	err = delKey(rdbFavorite, aHashId+LikeCountSuffix)
	return err
}

func (f Favorite) TruncateHateStatus(aHashId string) error {
	err := delKey(rdbFavorite, aHashId+hateSuffix)
	if err != nil {
		return err
	}
	err = delKey(rdbFavorite, aHashId+HateCountSuffix)
	return err
}
