package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/saul178/personal-website/middleware"
)

type RedisService struct {
	client *redis.Client
	ttl    time.Duration
}

func initNewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 3,
	})

	return client
}

func NewRedisCacheService() *RedisService {
	client := initNewRedisClient()
	ttl := time.Minute * 10

	return &RedisService{
		client: client,
		ttl:    ttl,
	}
}

// TODO: maybe have ttl dynamically? (ctx, key, val, ttl)
func (r *RedisService) Set(ctx context.Context, key string, val any) error {
	data, err := json.Marshal(val)
	if err != nil {
		middleware.ErrorLog.Printf("failed to marshall json into the cache: %v", err)
		return err
	}
	return r.client.Set(ctx, key, data, r.ttl).Err()
}

func (r *RedisService) Get(ctx context.Context, key string, dest any) error {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		middleware.ErrorLog.Printf("failed to Get cached data: %v", err)
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

func (r *RedisService) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
