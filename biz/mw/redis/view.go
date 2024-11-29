package redis

import "github.com/go-redis/redis/v7"

type View struct{}

const ViewCountSuffix = ":view_count"

func (c View) GetViewClient() *redis.Client {
	return rdbView
}

func (c View) IncrView(aHashId string) error {
	return incr(rdbView, aHashId+ViewCountSuffix)
}

func (c View) DecrView(aHashId string) error {
	return decr(rdbView, aHashId+ViewCountSuffix)
}

func (c View) CountView(aHashId string) (error, int64) {
	return nget(rdbView, aHashId+ViewCountSuffix)
}

func (c View) ViewCtAssign(aHashId string, count int64) error {
	return nset(rdbView, aHashId+ViewCountSuffix, count)
}

func (c View) CheckViewCt(aHashId string) (error, bool) {
	return check(rdbView, aHashId+ViewCountSuffix)
}

func (c View) DelViewCt(aHashId string) error {
	return delKey(rdbView, aHashId+ViewCountSuffix)
}

func (c View) GetViewMap(aHashIds []string) (error, map[string]int64) {
	return ngets(rdbView, aHashIds, ViewCountSuffix)
}
