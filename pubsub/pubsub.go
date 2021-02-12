package main

import (
	"fmt"
	pubsub2 "go-redis/pubsub"
)

func main() {
	pubsub := pubsub2.Subscribe()

	ch := pubsub.Channel()

	err := pubsub2.Publish()
	if err != nil {
		panic(err)
	}

	for msg := range ch{
		fmt.Println(msg.Channel, msg.Payload)
	}
}
