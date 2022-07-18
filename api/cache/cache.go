package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func setupClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func Set(ctx context.Context, key, value string) error {
	rc := setupClient()

	err := rc.Set(ctx, key, value, time.Second*30).Err()
	if err != nil {
		return err
	}
	return nil
}

func Get(ctx context.Context, key string) (string, error) {
	rc := setupClient()

	v, err := rc.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}

	return v, nil
}
