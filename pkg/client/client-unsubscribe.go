package client

import (
	"context"
	"github.com/spike-events/spike-events/bin"
	"log"
)

func (c *srv) unsubscribe(topic *bin.Topic) (*Success, error) {
	log.Println(">>> client unsubscribe: ", topic)
	success, err := c.client.Unsubscribe(context.Background(), topic)
	log.Println("<<< client unsubscribe: ", topic)
	if err != nil {
		return nil, err
	}
	return protoToSuccess(success), err
}
