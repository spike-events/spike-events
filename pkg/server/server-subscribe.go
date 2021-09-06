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
	nextSub *subscribe
	next    bool
}

func (s *server) registerSubscribe(topic *bin.Topic, eventServer bin.Spike_SubscribeServer) {
	defer s.m.Unlock()
	s.m.Lock()
	if s.subscribers == nil {
		s.subscribers = make(map[string][]*subscribe)
	}
	if topic.Persistent {
		dbTopic := s.registerSubscribeDatabase(topic)
		last := &subscribe{eventServer, topic, &dbTopic, nil, len(s.subscribers[topic.Topic]) == 0}
		s.subscribers[topic.Topic] = append(s.subscribers[topic.Topic], last)
		last.nextSub = s.subscribers[topic.Topic][len(s.subscribers[topic.Topic])-1]
		go func() {
			messages, err := s.db.TopicMessages(dbTopic, topic.GroupId, topic.Offset)
			if err != nil {
				log.Println(err)
			}
			for _, item := range messages {
				s.sendMessage(item, false)
			}
		}()
	} else {
		last := &subscribe{eventServer, topic, nil, nil, len(s.subscribersNonPersistence[topic.Topic]) == 0}
		s.subscribersNonPersistence[topic.Topic] = append(s.subscribersNonPersistence[topic.Topic], last)
		last.nextSub = s.subscribersNonPersistence[topic.Topic][len(s.subscribersNonPersistence[topic.Topic])-1]
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
