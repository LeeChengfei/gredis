package gredis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type Conf struct {
	Addr        string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	Wait        bool
}
type ReidsConn struct {
	redis.Conn
}

func (c *Conf) SetPool() redis.Pool {
	pool := redis.Pool{
		MaxIdle:     c.MaxIdle,
		MaxActive:   c.MaxActive,
		IdleTimeout: c.IdleTimeout,
		Wait:        c.Wait,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", c.Addr)
		},
	}

	return pool
}

func SetConn(pool redis.Pool) ReidsConn {
	defer pool.Close()
	conn := ReidsConn{
		pool.Get(),
	}
	return conn
}
