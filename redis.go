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

func (r *RedisAdapter[K]) Get(ctx context.Context, key K) (string, error) {
	return r.client.Get(ctx, string(key)).Result()
}

func (r *RedisAdapter[K]) GetInt64(ctx context.Context, key K) (int64, error) {
	return r.client.Get(ctx, string(key)).Int64()
}

func (r *RedisAdapter[K]) GetBool(ctx context.Context, key K) (bool, error) {
	return r.client.Get(ctx, string(key)).Bool()
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

func (r *RedisAdapter[K]) IncrBy(ctx context.Context, key K, val int64) error {
	return r.client.IncrBy(ctx, string(key), val).Err()
}

func (r *RedisAdapter[K]) Delete(ctx context.Context, key K) error {
	return r.client.Del(ctx, string(key)).Err()
}

func (r *RedisAdapter[K]) Close() error {
	return r.client.Close()
}
