package redis

import "github.com/go-redis/redis/v7"

type Comment struct{}

const CommentCountSuffix = ":comment_count"

func (c Comment) GetCommentClient() *redis.Client {
	return rdbComment
}

func (c Comment) IncrComment(aHashId string) error {
	return incr(rdbComment, aHashId+CommentCountSuffix)
}

func (c Comment) DecrComment(aHashId string) error {
	return decr(rdbComment, aHashId+CommentCountSuffix)
}

func (c Comment) CountComment(aHashId string) (error, int64) {
	return nget(rdbComment, aHashId+CommentCountSuffix)
}

func (c Comment) CommentCtAssign(aHashId string, count int64) error {
	return nset(rdbComment, aHashId+CommentCountSuffix, count)
}

func (c Comment) CheckCommentCt(aHashId string) (error, bool) {
	return check(rdbComment, aHashId+CommentCountSuffix)
}

func (c Comment) DelCommentCt(aHashId string) error {
	return delKey(rdbComment, aHashId+CommentCountSuffix)
}
