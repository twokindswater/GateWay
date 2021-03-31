package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

const noExpirationTime = 0

type redisDB struct {
	client *redis.Client
}

func InitRedis(addr string) *redisDB {
	return &redisDB{
		client: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
}

func (r *redisDB) Get(ctx context.Context, id string) ([]byte, error) {
	return r.client.Get(ctx, id).Bytes()
}

func (r *redisDB) Set(ctx context.Context, key string, data []byte) error {
	return r.client.Set(ctx, key, data, noExpirationTime).Err()
}
