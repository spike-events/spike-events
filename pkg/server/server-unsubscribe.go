package server

import (
	"context"
	"fmt"
	"github.com/spike-events/spike-events/bin"
)

func (s *server) Unsubscribe(ctx context.Context, topic *bin.Topic) (*bin.Success, error) {
	select {
	case <-ctx.Done():
		err := fmt.Errorf("context done")
		return responseError(err.Error()), err
	default:
	}
	err := ctx.Err()
	if err != nil {
		return responseError(err.Error()), err
	}
	s.unsubscribe(topic)
	return ResponseSuccess, nil
}

func (s *server) unsubscribe(topic *bin.Topic) {
	defer s.m.Unlock()
	s.m.Lock()

	if s.subscribers == nil {
		s.subscribers = make(map[string][]*subscribe)
	}
	var updatedSubscribers []*subscribe
	sub := s.subscribers[topic.Topic]
	for _, m := range sub {
		if m.topic.GetId() == topic.GetId() {
			m.msg <- &bin.Message{Offset: -1}
			<-m.success
			close(m.msg)
			close(m.success)
			continue
		}
		updatedSubscribers = append(updatedSubscribers, m)
	}
	s.subscribers[topic.GetTopic()] = updatedSubscribers
	if len(s.subscribers[topic.GetTopic()]) == 0 {
		delete(s.subscribers, topic.GetTopic())
	}

	if s.subscribersNonPersistence == nil {
		s.subscribersNonPersistence = make(map[string][]*subscribe)
	}

	var updatedSubscribersNonPersistent []*subscribe
	subNonPersistent := s.subscribersNonPersistence[topic.Topic]
	for _, m := range subNonPersistent {
		if m.topic.GetId() != topic.GetId() {
			m.msg <- &bin.Message{Offset: -1}
			<-m.success
			close(m.msg)
			close(m.success)
			continue
		}
		updatedSubscribersNonPersistent = append(updatedSubscribersNonPersistent, m)
	}
	s.subscribersNonPersistence[topic.GetTopic()] = updatedSubscribersNonPersistent
	if len(s.subscribersNonPersistence[topic.GetTopic()]) == 0 {
		delete(s.subscribersNonPersistence, topic.GetTopic())
	}
}
