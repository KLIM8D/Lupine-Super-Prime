package utils

import (
	"github.com/garyburd/redigo/redis"
	"runtime"
)

func NewFactory(host string) *RedisConf {
	v := &RedisConf{
		Host: host,
	}
	v.Pool = v.NewPool()

	return v
}

func (conf *RedisConf) NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, //Maximum number of connections
		Dial: func() (redis.Conn, error) {
			if c, err := redis.Dial("tcp", conf.Host); err != nil {
				panic(err.Error())
			} else {
				return c, nil
			}
		},
	}
}

func (conf *RedisConf) Add(item *RedisItem) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			err = r.(error)
		}
	}()

	c := conf.Pool.Get()
	defer c.Close()

	c.Send("SET", item.Key, item)

	if _, err := c.Do("EXEC"); err != nil {
		panic(err.Error())
	}

	return nil
}

func (conf *RedisConf) Get(key string, e interface{}) (interface{}, error) {
	c := conf.Pool.Get()
	defer c.Close()

	if r, err := c.Do("GET", key); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

func (conf *RedisConf) Exists(key string) (bool, error) {
	c := conf.Pool.Get()
	defer c.Close()

	if r, err := c.Do("EXISTS", key); err != nil {
		return false, err
	} else {
		return r.(int64) == 1, nil
	}
}
