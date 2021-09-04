package client

import (
	"context"
	"fmt"
)

type Subscriber struct {
	m       chan *Message
	connect bool
	topic   Topic
	ctx     context.Context
	client  *srv
}

func (s *Subscriber) Connected() bool {
	return s.connect
}

func (s *Subscriber) Close() error {
	_, err := s.client.Unsubscribe(s.topic)
	return err
}

func (s *Subscriber) Event() chan *Message {
	return s.m
}

func (s *Subscriber) ReadMessage() (*Message, error) {
	if s.connect {
		select {
		case <-s.ctx.Done():
			return nil, fmt.Errorf("context closed")
		case msg := <-s.m:
			return msg, nil
		}
	}
	return nil, fmt.Errorf("subscriber closed")
}
