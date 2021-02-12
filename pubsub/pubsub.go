package pubsub

import (
	"fmt"
)

func Pubsub() {
	pubsub := Subscribe()

	ch := pubsub.Channel()

	err := Publish()
	if err != nil {
		panic(err)
	}

	for msg := range ch{
		fmt.Println(msg.Channel, msg.Payload)
	}
}
