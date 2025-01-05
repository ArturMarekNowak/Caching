package caching

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

func OpenRedisConnection() (*redis.Client, context.Context) {
	ctx := context.Background()
	r := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS"),
		Password: "",
		DB:       0,
	})

	return r, ctx
}

func SetKey[T any](key string, str T) error {
	expirationDuration, err := time.ParseDuration(os.Getenv("CACHE_KEY_EXPIRATION"))
	if err != nil {
		return err
	}
	redisClient, ctx := OpenRedisConnection()
	err = redisClient.Set(ctx, key, str, expirationDuration).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetKey[T any](key string, value T) error {
	redisClient, ctx := OpenRedisConnection()
	err := redisClient.Get(ctx, key).Scan(value)
	if err != nil {
		return err
	}
	return nil
}

func KeyExists[T any](key string, value T) bool {
	return GetKey(key, value) != nil
}

func DelKey(key string) error {
	redisClient, ctx := OpenRedisConnection()
	err := redisClient.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
