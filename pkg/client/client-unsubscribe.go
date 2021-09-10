package client

import (
	"context"
	"github.com/spike-events/spike-events/bin"
)

func (c *srv) unsubscribe(topic *bin.Topic) (*Success, error) {
	success, err := c.client.Unsubscribe(context.Background(), topic)
	if err != nil {
		return nil, err
	}
	return protoToSuccess(success), err
}
