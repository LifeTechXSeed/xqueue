package redis

import (
	"xqueue/log"
	"xqueue/util"
)

var logger = log.NewLogger("redis")

type RedisCli interface {
	Close() error
}

func CreateNewCache() RedisCli {
	redisHost := util.GetEnv("REDIS_URI", "127.0.0.1")
	redisPort := util.GetEnv("REDIS_PORT", "6379")

	redis := NewRedisClient(redisHost, redisPort)

	return redis
}
