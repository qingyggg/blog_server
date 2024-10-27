package redis

import (
	"errors"
	"github.com/go-redis/redis/v7"
)

// redis set
// add k & v to redis
func sadd(c *redis.Client, k string, v interface{}) error {
	tx := c.TxPipeline()
	tx.SAdd(k, v)
	tx.Expire(k, ExpireTime)
	_, err := tx.Exec()
	return err
}

// srem k & v
func srem(c *redis.Client, k string, v interface{}) error {
	tx := c.TxPipeline()
	tx.SRem(k, v)
	_, err := tx.Exec()
	return err
}

// exist check the relation k and v if exist
func sexist(c *redis.Client, k string, v string) (error, bool) {
	e, err := c.SIsMember(k, v).Result()
	if err != nil {
		return err, false
	}
	if e {
		c.Expire(k, ExpireTime)
		return nil, true
	}
	return nil, false
}

// scount sget the size of the set of key
func scount(c *redis.Client, k string) (err error, sum int64) {
	if sum, err = c.SCard(k).Result(); err == nil {
		c.Expire(k, ExpireTime)
		return err, sum
	}
	return err, sum
}

func sget(c *redis.Client, k string) (error, *[]string) {
	var vt []string
	v, err := c.SMembers(k).Result()
	if err != nil {
		return err, nil
	}
	c.Expire(k, ExpireTime)
	for _, vs := range v {
		vt = append(vt, vs)
	}
	return nil, &vt
}

// redis common
// delKey 删除键 k 及其关联的值
func delKey(c *redis.Client, k string) error {
	// 开启 Redis 事务
	tx := c.TxPipeline()
	// 使用 DEL 命令删除键 k
	tx.Del(k)
	// 执行事务
	_, err := tx.Exec()
	// 返回执行的错误信息
	return err
}

// check the set of k if exist
func check(c *redis.Client, k string) (error, bool) {
	e, err := c.Exists(k).Result()

	if err != nil {
		return err, false
	}
	if e > 0 {
		return nil, true
	}
	return nil, false
}

// redis num
func nset(c *redis.Client, k string, v int64) error {
	tx := c.TxPipeline()
	tx.Set(k, v, ExpireTime)
	_, err := tx.Exec()
	return err
}
func incr(c *redis.Client, k string) error {
	tx := c.TxPipeline()
	tx.Incr(k)
	tx.Expire(k, ExpireTime)
	_, err := tx.Exec()
	return err
}
func decr(c *redis.Client, k string) error {
	tx := c.TxPipeline()
	tx.Decr(k)
	tx.Expire(k, ExpireTime)
	_, err := tx.Exec()
	return err
}

// 搭配incr,decr,用以获取(评论数，点赞数)统计的数量
func nget(c *redis.Client, k string) (error, int64) {
	// 使用 GET 命令获取键对应的值
	val, err := c.Get(k).Int64() // 将结果转换为 int64 类型
	if err != nil {
		return err, 0 // 如果有错误，返回 0 和错误信息
	}
	c.Expire(k, ExpireTime)
	return nil, val // 返回获取到的值
}

// 如果查询的键不存在，默认赋值为-1
func ngets(c *redis.Client, ks []string, suffix string) (error, map[string]int64) {
	resMap := make(map[string]int64)

	// 开始管道操作
	pipe := c.Pipeline()
	cmds := make([]*redis.StringCmd, len(ks))

	// 添加 GET 命令到管道
	for i, k := range ks {
		cmds[i] = pipe.Get(k + suffix) // 传入上下文
	}

	// 执行管道操作
	if _, err := pipe.Exec(); err != nil && !errors.Is(err, redis.Nil) {
		return err, nil
	}

	// 获取结果
	for i, cmd := range cmds {
		value, err := cmd.Int64()

		if err != nil {
			if errors.Is(err, redis.Nil) {
				// 如果键不存在，可以选择跳过或记录信息
				resMap[ks[i]] = -1
			} else {
				return err, nil // 处理其他错误
			}
		}

		resMap[ks[i]] = value
	}
	return nil, resMap
}
