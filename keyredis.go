package gredis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

//检查key是否存在
func (r ReidsConn) Exists(key string) (bool, error) {
	return redis.Bool(r.Do("exists", key))
}

//删除key
func (r ReidsConn) Del(key string) error {
	_, err := r.Do("del", key)
	return err
}

//设置key的过期时间
func (r ReidsConn) Expire(key string, second time.Duration) error {
	_, err := r.Do("expire", key, second)
	return err
}

//重命名key
func (r ReidsConn) Rename(key, newkey string) error {
	_, err := r.Do("rename", key, newkey)
	return err
}
