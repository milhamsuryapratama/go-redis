package pubsub

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-redis/pubsub/config"
)

func Subscribe() *redis.PubSub {
	var (
		rdb = config.RedisConfig()
		ctx = context.Background()
	)

	pubsub := rdb.Subscribe(ctx, "mychannel1")

	_, err := pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	}

	return pubsub
}