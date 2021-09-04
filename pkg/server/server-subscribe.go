package server

import (
	"fmt"
	"log"
	"spike.io/bin"
	"spike.io/internal/models"
)

type subscribe struct {
	bin.Spike_SubscribeServer
	topic   *bin.Topic
	dbTopic *models.Topic
}

func (s *server) registerSubscribe(topic *bin.Topic, eventServer bin.Spike_SubscribeServer) {
	defer s.m.Unlock()
	s.m.Lock()
	if s.subscribers == nil {
		s.subscribers = make(map[string][]*subscribe)
	}
	if topic.Persistent {
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
	} else {
		s.subscribersNonPersistence[topic.Topic] = append(s.subscribersNonPersistence[topic.Topic], &subscribe{eventServer, topic, nil})
	}
}

func (s *server) registerSubscribeDatabase(topic *bin.Topic) models.Topic {
	dbTopic, err := s.db.Subscribe(topic.Topic, topic.GroupId, topic.Offset)
	if err != nil {
		log.Println(err)
	}
	return dbTopic
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
