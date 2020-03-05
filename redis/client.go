package redis

import (
	"github.com/go-redis/redis"
)

type RedisClient struct {
	client *redis.Client
}

func (r *RedisClient) Close() error {
	return r.client.Close()
}

func NewRedisClient(host, port string) *RedisClient {
	cli := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
	})
	return &RedisClient{
		client: cli,
	}
}
