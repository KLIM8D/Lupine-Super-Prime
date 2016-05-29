package utils

import (
	"github.com/garyburd/redigo/redis"
)

type RedisConf struct {
	Pool *redis.Pool
	Host string
}

type RedisItem struct {
	Key   uint64
	Value interface{}
}

type SchedulerConf struct {
	Amount int
	Master bool
}

type Configuration struct {
	ConfigPath string
	Redis      *RedisConf
	Scheduler  *SchedulerConf
}
