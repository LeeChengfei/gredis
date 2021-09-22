package gredis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type Conf struct {
	Addr        string        //连接地址
	MaxIdle     int           //连接池最大空闲连接数
	MaxActive   int           //连接池在给定时间内分配的最大连接数。 当为0时，池中的连接数没有限制
	IdleTimeout time.Duration //连接超时
	Wait        bool          //如果Wait为真，并且连接池处于MaxActive限制，那么Get()将等待连接返回到池中，然后返回。
	DB          string        //数据库
	Password    string        //密码
}
type ReidsConn struct {
	redis.Conn
}

//建立连接池
func (c *Conf) SetPool() redis.Pool {
	pool := redis.Pool{
		MaxIdle:     c.MaxIdle,
		MaxActive:   c.MaxActive,
		IdleTimeout: c.IdleTimeout,
		Wait:        c.Wait,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", c.Addr)
			if err != nil {
				return conn, err
			}
			//如果密码不为空，输入密码
			if c.Password != "" {
				_, err = conn.Do("auth", c.Password)
				if err != nil {
					return conn, err
				}
			}
			//选择数据库
			if c.DB != "" {
				_, err = conn.Do("select", c.DB)
				if err != nil {
					return conn, err
				}
			}
			return conn, err
		},
	}

	return pool
}

//获取链接
func SetConn(pool redis.Pool) ReidsConn {
	defer pool.Close()
	conn := ReidsConn{
		pool.Get(),
	}
	return conn
}
