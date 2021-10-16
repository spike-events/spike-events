package client

import (
	"context"
	"log"
)

func (c *srv) Publish(message Message) (*Success, error) {
	log.Println(">>> client publish: ", message)
	success, err := c.client.Publish(context.Background(), message.ProtoMessage())
	log.Println("<<< client publish: ", message)
	if err != nil {
		return nil, err
	}
	return protoToSuccess(success), err
}
