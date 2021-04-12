package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

const noExpirationTime = 0

type redisDB struct {
	client *redis.Client
}

func InitRedis(addr string) *redisDB {
	return &redisDB{
		client: redis.NewClient(&redis.Options{
			Addr:               addr,
			IdleTimeout:        time.Second * 30,
			IdleCheckFrequency: time.Second * 5,
		}),
	}
}

func (r *redisDB) Get(ctx context.Context, id string) ([]byte, error) {
	cmd := r.client.Get(ctx, id)
	if len(cmd.Val()) == 0 {
		return nil, nil
	} else {
		return cmd.Bytes()
	}
}

func (r *redisDB) Set(ctx context.Context, key string, data []byte) error {
	return r.client.Set(ctx, key, data, noExpirationTime).Err()

}
