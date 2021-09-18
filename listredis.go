package gredis

import "github.com/gomodule/redigo/redis"

//从左边插入value
func (r ReidsConn) LPush(key string, value ...interface{}) error {
	_, err := r.Do("lpush", key, value)
	return err
}

//从右边插入value
func (r ReidsConn) RPush(key string, value ...interface{}) error {
	_, err := r.Do("rpush", key, value)
	return err
}

//从左边移除第一个元素
func (r ReidsConn) LPop(key string) (string, error) {
	return redis.String(r.Do("lpop", key))
}

//从右边移除第一个元素
func (r ReidsConn) RPop(key string) (string, error) {
	return redis.String(r.Do("rpop", key))
}

//根据索引index获取列表元素
func (r ReidsConn) LIndex(key string, index int) (string, error) {
	return redis.String(r.Do("lindex", key, index))
}

//获取指定范围内的所有元素
func (r ReidsConn) LRange(key string, start, end int) ([]string, error) {
	return redis.Strings(r.Do("lrange", key, start, end))
}

//在指定元素pitot前面插入value
func (r ReidsConn) LInsertBefore(key string, value, pirot interface{}) error {
	_, err := r.Do("linsert", key, "before", pirot, value)
	return err
}

//在指定元素pitot后面插入value
func (r ReidsConn) LInsertAfter(key string, value, pirot interface{}) error {
	_, err := r.Do("linsert", key, "after", pirot, value)
	return err
}

//删除list中count个的value元素
//当count<0时，从右到左开始，删除|count|个value元素
//当count>0时，从左到右开始，删除count个value元素
//放count=0时，删除lint中所有的value元素
func (r ReidsConn) LRem(key string, value interface{}, count int) error {
	_, err := r.Do("lrem", key, count, value)
	return err
}

//删除区间之外的所有元素
func (r ReidsConn) LTrim(key string, start, end int) error {
	_, err := r.Do("ltrim", key, start, end)
	return err
}
