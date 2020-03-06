package redis

import (
	"xqueue/log"
	"xqueue/util"

	"github.com/go-redis/redis"
)

var logger = log.NewLogger("redis")

type RedisCli interface {
	Close() error
	Subscribe(channel string) *redis.PubSub
	Publish(chanel, message string) error
	ZRange(key string, start, stop int64) ([]string, error)
	ZJobToQueue(key string, job_id, priority int) error
	HSet(key, field, value string) error
	HGet(key, field string) (string, error)
}

func CreateNewCache() RedisCli {
	redisHost := util.GetEnv("REDIS_URI", "127.0.0.1")
	redisPort := util.GetEnv("REDIS_PORT", "6379")

	redis := NewRedisClient(redisHost, redisPort)

	return redis
}
