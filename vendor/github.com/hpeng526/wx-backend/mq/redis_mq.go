package mq

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisMq struct {
	Pool *redis.Pool
}

func NewRedisMq(server string) *RedisMq {
	Pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return &RedisMq{Pool}
}

func (mq *RedisMq) Poll(key string, duration time.Duration) (value string) {
	conn := mq.getConn()
	defer conn.Close()
	reply, err := conn.Do("BLPOP", key, duration.Seconds())
	if err != nil {
		return
	}
	ans, err := redis.Strings(reply, err)
	if err != nil {
		return
	}
	value = ans[1]
	return
}

func (mq *RedisMq) Offer(key string, value string) (err error) {
	conn := mq.getConn()
	defer conn.Close()
	_, err = redis.Int64(conn.Do("RPUSH", key, value))
	return
}

func (mq *RedisMq) getConn() redis.Conn {
	conn := mq.Pool.Get()
	if conn.Err() != nil {
		panic(conn.Err().Error())
	}
	return conn
}
