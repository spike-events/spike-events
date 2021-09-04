package server

import (
	"context"
	"fmt"
	"spike.io/bin"
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

	if s.monitors == nil {
		s.monitors = make(map[string][]*monitor)
	}
	var updatedMonitors []*monitor
	mon := s.monitors[topic.Topic]
	for _, m := range mon {
		if m.topic.GetGroupId() != topic.GetGroupId() {
			updatedMonitors = append(updatedMonitors, m)
		} else {
			m.Spike_MonitorEventServer.Send(&bin.Message{Offset: -1})
		}
	}
	s.monitors[topic.GetTopic()] = updatedMonitors
	if len(s.monitors[topic.GetTopic()]) == 0 {
		delete(s.monitors, topic.GetTopic())
	}

	if s.subscribers == nil {
		s.subscribers = make(map[string][]*subscribe)
	}
	var updatedSubscribers []*subscribe
	sub := s.subscribers[topic.Topic]
	for _, m := range sub {
		if m.topic.GetGroupId() != topic.GetGroupId() {
			updatedSubscribers = append(updatedSubscribers, m)
		} else {
			m.Spike_SubscribeServer.Send(&bin.Message{Offset: -1})
		}
	}
	s.subscribers[topic.GetTopic()] = updatedSubscribers
	if len(s.subscribers[topic.GetTopic()]) == 0 {
		delete(s.subscribers, topic.GetTopic())
	}
}

