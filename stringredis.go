package gredis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

//设置字符串key的值
func (r ReidsConn) Set(key string, value interface{}) error {
	_, err := r.Do("set", key, value)
	return err
}

//获取字符串的值
func (r ReidsConn) Get(key string) (string, error) {
	return redis.String(r.Do("get", key))
}

//设置key的新值，同时返回key的旧值
func (r ReidsConn) GetSet(key string, value interface{}) (string, error) {
	return redis.String(r.Do("getset", key, value))
}

//往key追加值
func (r ReidsConn) Append(key string, value interface{}) error {
	_, err := r.Do("append", key, value)
	return err
}

//获取字符串的长度
func (r ReidsConn) StrLen(key string) (int64, error) {
	return redis.Int64(r.Do("strlen", key))
}

//设置字符擦的值，同时设置过期时间
func (r ReidsConn) Setex(key string, value interface{}, second time.Duration) error {
	_, err := r.Do("setex", key, second, value)
	return err
}

//值加num
func (r ReidsConn) Incrby(key string, num int) error {
	_, err := r.Do("incrby", key, num)
	return err
}

//值减num
func (r ReidsConn) Decrby(key string, num int) error {
	_, err := r.Do("decrby", key, num)
	return err
}
