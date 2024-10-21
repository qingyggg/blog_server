package redis

import "github.com/go-redis/redis/v7"

type Collect struct{}

const CollectCountSuffix = ":collect_count"

func (c Collect) GetCollectClient() *redis.Client {
	return rdbCollect
}

func (c Collect) IncrCollect(aHashId string) error {
	return incr(rdbCollect, aHashId+CollectCountSuffix)
}

func (c Collect) DecrCollect(aHashId string) error {
	return decr(rdbCollect, aHashId+CollectCountSuffix)
}

func (c Collect) CountCollect(aHashId string) (error, int64) {
	return nget(rdbCollect, aHashId+CollectCountSuffix)
}

func (c Collect) CollectCtAssign(aHashId string, count int64) error {
	return nset(rdbCollect, aHashId+CollectCountSuffix, count)
}

func (c Collect) CheckCollectCt(aHashId string) (error, bool) {
	return check(rdbCollect, aHashId+CollectCountSuffix)
}

func (c Collect) DelCollectCt(aHashId string) error {
	return delKey(rdbCollect, aHashId+CollectCountSuffix)
}
