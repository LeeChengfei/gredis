package gredis

import "github.com/gomodule/redigo/redis"

//想集合key添加成员
func (r ReidsConn) SAdd(key string, value ...interface{}) error {
	_, err := r.Do("sadd", key, value)
	return err
}

//获取结合key的成员总数
func (r ReidsConn) SCard(key string) (int, error) {
	return redis.Int(r.Do("scard", key))
}

//集合key与其他集合others相较，获取集合key独有的元素
func (r ReidsConn) SDiff(key string, others ...string) ([]string, error) {

	return redis.Strings(r.Do("sdiff", key, others))
}

//返回所有集合的交集
func (r ReidsConn) SInter(key string, others ...string) ([]string, error) {
	return redis.Strings(r.Do("sinter", key, others))
}

//显示集合所有成员
func (r ReidsConn) SMembers(key string) ([]string, error) {
	return redis.Strings(r.Do("smembers", key))
}

//判断member是否为key的成员
func (r ReidsConn) SIsmenber(key string, member interface{}) (bool, error) {
	return redis.Bool(r.Do("sismember", key, member))
}

//移除集合count个成员，并将其返回
func (r ReidsConn) SPop(key string, count int) ([]string, error) {
	return redis.Strings(r.Do("spop", key, count))
}

//将key1的成员member移动到key2中
func (r ReidsConn) SMove(key1, key2 string, member interface{}) error {
	_, err := r.Do("smove", key1, key2, member)
	return err
}

//删除集合成员
func (r ReidsConn) SRem(key string, members ...interface{}) error {
	_, err := r.Do("srem", key, members)
	return err
}

//返回所有集合的并集
func (r ReidsConn) SUnion(key ...string) ([]string, error) {
	return redis.Strings(r.Do("sunion", key))
}
