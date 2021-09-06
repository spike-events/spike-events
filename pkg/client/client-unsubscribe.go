package client

import (
	"context"
)

func (c *srv) Unsubscribe(topic Topic) (*Success, error) {
	success, err := c.client.Unsubscribe(context.Background(), topic.ProtoMessage())
	if err != nil {
		return nil, err
	}
	return protoToSuccess(success), err
}
