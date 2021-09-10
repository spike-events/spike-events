package client

import (
	"context"
	"github.com/spike-events/spike-events/bin"
	"sync"
)

type Subscriber struct {
	m       chan *Message
	connect bool
	topic   *bin.Topic
	ctx     context.Context
	client  *srv
	wg      sync.WaitGroup
}

func (s *Subscriber) Connected() bool {
	return s.connect
}

func (s *Subscriber) Close() error {
	_, err := s.client.unsubscribe(s.topic)
	s.wg.Wait()
	return err
}

func (s *Subscriber) Event() chan *Message {
	return s.m
}
