package cache

import (
	"context"

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

	err := rc.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
