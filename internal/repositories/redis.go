package repositories

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(url string) *RedisClient {
	opt, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}
	return &RedisClient{
		client: redis.NewClient(opt),
	}
}

func (rc *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return rc.client.Get(ctx, key).Result()
}

func (rc *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return rc.client.Set(ctx, key, value, expiration).Err()
}

func (rc *RedisClient) Delete(ctx context.Context, key string) error {
	return rc.client.Del(ctx, key).Err()
}

func (rc *RedisClient) DeleteByPattern(ctx context.Context, pattern string) error {
	keys, err := rc.client.Keys(ctx, pattern).Result()
	if err != nil {
		return err
	}
	return rc.client.Del(ctx, keys...).Err()
}

func (rc *RedisClient) Close() {
	rc.client.Close()
}
