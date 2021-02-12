package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

func RedisConfig() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	return rdb
}

func Set(ctx context.Context, rdb *redis.Client, key string, data interface{}, expired time.Duration) error {
	err := rdb.Set(ctx, key, data, expired).Err()
	return err
}

func Get(ctx context.Context, rdb *redis.Client, key string) (val string, err error) {
	val, err = rdb.Get(ctx, key).Result()
	return
}