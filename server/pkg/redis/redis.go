package redis

import "github.com/redis/go-redis/v9"

func New() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis_server:6379",
	})
	return rdb
}
