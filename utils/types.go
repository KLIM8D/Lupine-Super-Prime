package utils

import (
	"github.com/garyburd/redigo/redis"
)

type RedisConf struct {
	Pool *redis.Pool
	Host string
}

type RedisItem struct {
	Key   float64
	Value interface{}
}

type Configuration struct {
	ConfigPath string
	Redis      RedisConf
}
