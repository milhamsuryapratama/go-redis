package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

type person struct {
	Name, Address string
}

func main() {

	p := person{
		Name: "ILham",
		Address: "Maron",
	}

	rdb := setupRedis()

	data, _ := json.Marshal(p)

	// set to 0 jika kamu tidak ingin memberikan waktu expired pada redis
	err := Set(ctx, rdb, "person", data, 0)

	if err != nil {
		panic(err)
	}

	val, err := Get(ctx, rdb, "person")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(val), &p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)
}

func setupRedis() *redis.Client {
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
