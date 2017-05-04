package cache

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisCache struct {
	Pool *redis.Pool
}

func NewRedisCache(server string) *RedisCache {
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
	return &RedisCache{Pool}
}

func (rc *RedisCache) Get(key string) interface{} {
	conn := rc.getConn()
	defer conn.Close()
	var data string
	data, err := redis.String(conn.Do("GET", key))
	if err == nil {
		return data
	}
	return nil
}

func (rc *RedisCache) Set(key string, value interface{}, timeout time.Duration) error {
	conn := rc.getConn()
	defer conn.Close()
	var s string
	if v, ok := value.(string); ok {
		s = v
	}
	_, err := conn.Do("SET", key, s, "EX", timeout.Seconds())
	return err
}

func (rc *RedisCache) IsExist(key string) bool {
	conn := rc.getConn()
	defer conn.Close()
	data, err := redis.Bool(conn.Do("EXISTS", key))
	if err == nil {
		return data
	}
	return false
}

func (rc *RedisCache) Delete(key string) error {
	conn := rc.getConn()
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	return err
}

func (rc *RedisCache) getConn() redis.Conn {
	conn := rc.Pool.Get()
	if conn.Err() != nil {
		panic(conn.Err().Error())
	}
	return conn
}
