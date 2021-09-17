package gredis

import "github.com/gomodule/redigo/redis"

//建立哈希表自字段
func (r ReidsConn) HSet(key string, field string, value interface{}) error {
	_, err := r.Do("hset", key, field, value)
	return err
}

//获取哈希字段的值
func (r ReidsConn) HGet(key, field string) (string, error) {
	return redis.String(r.Do("hget", key, field))
}

//获取所有的哈希字段的值
func (r ReidsConn) HGetall(key string) (map[string]string, error) {

	mm, err := redis.StringMap(r.Do("hgetall", key))
	if err != nil {
		return nil, err
	}
	return mm, err
}

//同时设置多个哈希字段
func (r ReidsConn) HMset(key string, field map[string]string) error {
	_, err := r.Do("hmset", key, field)
	return err
}

//获取哈希字段
func (r ReidsConn) HKeys(key string) ([]string, error) {
	return redis.Strings(r.Do("hkeys", key))
}

//获取哈希表所有的值
func (r ReidsConn) HVals(key string) ([]string, error) {
	return redis.Strings(r.Do("hvals", key))
}

//哈希表字段的值加num
func (r ReidsConn) HIncrby(key, field string, num int) error {
	_, err := r.Do("hincrby", key, field, num)
	return err
}

//哈希表字段的值减num
func (r ReidsConn) HDecrby(key, field string, num int) error {
	_, err := r.Do("hdecrby", key, field, num)
	return err
}

//哈希表字段数量
func (r ReidsConn) HLen(key string) (int64, error) {
	return redis.Int64(r.Do("hlen", key))
}

//删除哈希字段
func (r ReidsConn) HDel(key string, field string) error {
	_, err := r.Do("hdel", key, field)
	return err
}
