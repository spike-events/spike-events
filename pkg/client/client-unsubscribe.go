package client

import (
	"context"
	"spike.io/bin"
)

func (c *srv) unsubscribe(topic *bin.Topic) (*Success, error) {
	success, err := c.client.Unsubscribe(context.Background(), topic)
	if err != nil {
		return nil, err
	}
	return protoToSuccess(success), err
}
