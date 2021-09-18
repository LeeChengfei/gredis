package gredis

import "github.com/gomodule/redigo/redis"

//设置有序集合的值
func (r ReidsConn) ZAdd(key string, score int, member interface{}) error {
	_, err := r.Do("zadd", key, score, member)
	return err
}

//获取有序集合的成员数
func (r ReidsConn) ZCard(key string) (int, error) {
	return redis.Int(r.Do("zcard", key))
}

//计算min，max区间中所有成员的总分数
func (r ReidsConn) ZCount(key string, min, max int) (int, error) {
	return redis.Int(r.Do("zcount", key, min, max))
}

//有序集合key的member成员分数加incrment
func (r ReidsConn) ZIncrby(key string, incrememt int, member interface{}) (int, error) {
	return redis.Int(r.Do("zincrby", key, incrememt, member))
}

//有序集合key的member成员分数减incrment
func (r ReidsConn) ZDecrby(key string, incrememt int, member interface{}) (int, error) {
	return redis.Int(r.Do("decrby", key, incrememt, member))
}

//获取所有成员的总分数
func (r ReidsConn) ZScore(key string, member interface{}) (int, error) {
	return redis.Int(r.Do("zscore", key, member))
}

//获取索引start，end区间内的所有成员
func (r ReidsConn) ZRange(key string, start, end int) ([]string, error) {
	return redis.Strings(r.Do("zrange", key, start, end))
}

//获取指定member成员的索引
func (r ReidsConn) ZRank(key string, member interface{}) (int, error) {
	return redis.Int(r.Do("zrank", key, member))
}

//删除成员
func (r ReidsConn) ZRem(key string, members ...interface{}) error {
	_, err := r.Do("zrem", key, members)
	return err
}
