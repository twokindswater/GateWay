package redis

import "github.com/go-redis/redis/v8"

func InitRedis(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}
