package client

import (
	"context"
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
