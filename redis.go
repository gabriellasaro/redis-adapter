package redisadapter

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisAdapter[K ~string] struct {
	client *redis.Client
}

func NewARedis[K ~string](client *redis.Client) *RedisAdapter[K] {
	return &RedisAdapter[K]{client}
}

func (r *RedisAdapter[K]) GetJSON(ctx context.Context, key K, dest any) error {
	data, err := r.client.Get(ctx, string(key)).Bytes()
	if err != nil {
		return err
	}

	return json.Unmarshal(data, dest)
}

func (r *RedisAdapter[K]) SetJSON(ctx context.Context, key K, value any, exp time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, string(key), v, exp).Err()
}
