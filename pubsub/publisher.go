package pubsub

import (
	"context"
	"go-redis/pubsub/config"
)

func Publish() error {
	var (
		rdb = config.RedisConfig()
		ctx = context.Background()
	)

	err := rdb.Publish(ctx, "mychannel1", "yolo").Err()
	if err != nil {
		panic(err)
	}

	return err
}