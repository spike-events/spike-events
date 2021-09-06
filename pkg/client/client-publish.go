package client

import (
	"context"
)

func (c *srv) Publish(message Message) (*Success, error) {
	success, err := c.client.Publish(context.Background(), message.ProtoMessage())
	if err != nil {
		return nil, err
	}
	return protoToSuccess(success), err
}
