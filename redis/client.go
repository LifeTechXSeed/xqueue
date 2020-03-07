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

func (r *RedisClient) Del(key string) error {
	err := r.client.Del(key).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisClient) Subscribe(channel string) *redis.PubSub {
	return r.client.Subscribe(channel)
}

func (r *RedisClient) Publish(chanel, message string) error {
	err := r.client.Publish(chanel, message).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisClient) ZJobToQueue(key string, job_id, priority int) error {
	mem := redis.Z{Score: float64(priority), Member: job_id}
	err := r.client.ZAdd(key, mem).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *RedisClient) ZRange(key string, start, stop int64) ([]string, error) {
	result, err := r.client.ZRange(key, start, stop).Result()
	if err == redis.Nil {
		return []string{}, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *RedisClient) ZRem(key, id string) error {
	err := r.client.ZRem(key, id).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) ZCount(key, max, min string) (int64, error) {
	number, err := r.client.ZCount(key, max, min).Result()
	if err != nil {
		return -1, err
	}

	return number, nil
}

func (r *RedisClient) HSet(key, field, value string) error {
	err := r.client.HSet(key, field, value).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisClient) HGet(key, field string) (string, error) {
	result, err := r.client.HGet(key, field).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		logger.Error(err)
		return "", err
	}

	return result, nil
}
