package server

import (
	"fmt"
	"log"
	"spike.io/bin"
	"spike.io/internal/models"
)

type monitor struct {
	topic *bin.Topic
	bin.Spike_MonitorEventServer
}

type subscribe struct {
	bin.Spike_SubscribeServer
	topic   *bin.Topic
	dbTopic *models.Topic
}

func (s *server) registerMonitor(topic *bin.Topic, eventServer bin.Spike_MonitorEventServer) {
	defer s.m.Unlock()
	s.m.Lock()
	if s.monitors == nil {
		s.monitors = make(map[string][]*monitor)
	}
	s.monitors[topic.Topic] = append(s.monitors[topic.Topic], &monitor{topic, eventServer})
}

func (s *server) registerSubscribe(topic *bin.Topic, eventServer bin.Spike_SubscribeServer) {
	defer s.m.Unlock()
	s.m.Lock()
	if s.subscribers == nil {
		s.subscribers = make(map[string][]*subscribe)
	}
	dbTopic := s.registerSubscribeDatabase(topic)
	s.subscribers[topic.Topic] = append(s.subscribers[topic.Topic], &subscribe{eventServer, topic, &dbTopic})
	go func() {
		messages, err := s.db.TopicMessages(dbTopic, topic.Topic, topic.GroupId, topic.Offset)
		if err != nil {
			log.Println(err)
		}
		for _, item := range messages {
			s.sendMessage(item, false)
		}
	}()
}

func (s *server) registerSubscribeDatabase(topic *bin.Topic) models.Topic {
	dbTopic, err := s.db.Subscribe(topic.Topic, topic.GroupId, topic.Offset)
	if err != nil {
		log.Println(err)
	}
	return dbTopic
}

func (s *server) MonitorEvent(topic *bin.Topic, eventServer bin.Spike_MonitorEventServer) error {
	select {
	case <-eventServer.Context().Done():
		return fmt.Errorf("closed monitor event")
	default:
	}
	s.registerMonitor(topic, eventServer)
	return nil
}

func (s *server) Subscribe(topic *bin.Topic, subscribeServer bin.Spike_SubscribeServer) error {
	select {
	case <-subscribeServer.Context().Done():
		return fmt.Errorf("closed monitor event")
	default:
	}
	if len(topic.Topic) == 0 {
		return fmt.Errorf("topic is required")
	}
	if len(topic.GroupId) == 0 {
		return fmt.Errorf("groupId is required")
	}
	s.registerSubscribe(topic, subscribeServer)
	return nil
}
