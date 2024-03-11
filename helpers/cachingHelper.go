package helpers

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

func CreateClient() (*redis.Client, context.Context) {
	ctx := context.Background()
	r := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS"),
		Password: "",
		DB:       0,
	})

	return r, ctx
}

func SetKey[S any](key string, str S) error {
	expirationDuration, err := time.ParseDuration(os.Getenv("CACHE_KEY_EXPIRATION"))
	if err != nil {
		fmt.Printf("Couldn't parse CACHE_KEY_EXPIRATION: %s\n", expirationDuration)
		return err
	}
	redisClient, ctx := CreateClient()
	err = redisClient.Set(ctx, key, str, expirationDuration).Err()
	if err != nil {
		fmt.Printf("Couldn't save key: %s, error: %s\n", key, err)
		return err
	}
	return nil
}

func GetKey[S any](key string, str S) error {
	redisClient, ctx := CreateClient()
	err := redisClient.Get(ctx, key).Scan(str)
	if err != nil {
		fmt.Printf("Couldn't get key: %s, error: %s\n", key, err)
		return err
	}
	return nil
}

func DelKey(key string) error {
	redisClient, ctx := CreateClient()
	err := redisClient.Del(ctx, key).Err()
	if err != nil {
		fmt.Printf("Couldn't del key: %s, error: %s\n", key, err)
		return err
	}
	return nil
}
