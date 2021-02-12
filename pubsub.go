package main

import "fmt"

func main() {
	rdb := setupRedis()

	pubsub := rdb.Subscribe(ctx, "mychannel1")

	_, err := pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	}

	ch := pubsub.Channel()

	//err = rdb.Publish(ctx, "mychannel1", "ilham").Err()
	//if err != nil {
	//	panic(err)
	//}

	for msg := range ch{
		fmt.Println(msg.Channel, msg.Payload)
	}
}
