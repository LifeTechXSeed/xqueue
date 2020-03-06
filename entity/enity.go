package entity

import "xqueue/redis"

type Entity struct {
	Redis redis.RedisCli
}

func CreateNewEntity() *Entity {
	redisIns := redis.CreateNewCache()
	return &Entity{
		Redis: redisIns,
	}
}

func (e *Entity) Release() {
	e.Redis.Close()
}
