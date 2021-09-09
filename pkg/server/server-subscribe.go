package server

import (
	"fmt"
	"log"
	"spike.io/bin"
	"spike.io/internal/models"
)

type subscribe struct {
	id      string
	msg     chan *bin.Message
	success chan error
	topic   *bin.Topic
	dbTopic *models.Topic
	next    bool
}

func (s *server) registerSubscribe(topic *bin.Topic) (chan *bin.Message, chan error) {
	defer s.m.Unlock()
	s.m.Lock()
	if s.subscribers == nil {
		s.subscribers = make(map[string][]*subscribe)
	}
	channel := make(chan *bin.Message, 100)
	success := make(chan error, 100)
	if topic.Persistent {
		dbTopic := s.registerSubscribeDatabase(topic)
		last := &subscribe{topic.GetId(), channel, success, topic, &dbTopic, len(s.subscribers[topic.Topic]) == 0}
		s.subscribers[topic.Topic] = append(s.subscribers[topic.Topic], last)
		go func() {
			messages, err := s.db.TopicMessages(dbTopic, topic.GroupId, topic.Offset)
			if err != nil {
				log.Println(err)
			}
			for _, item := range messages {
				s.sendMessage(item, false, topic.GetId())
			}
		}()
	} else {
		last := &subscribe{topic.GetId(), channel, success, topic, nil, len(s.subscribersNonPersistence[topic.Topic]) == 0}
		s.subscribersNonPersistence[topic.Topic] = append(s.subscribersNonPersistence[topic.Topic], last)
	}
	return channel, success
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
	cMessages, cSuccess := s.registerSubscribe(topic)
	for msg := range cMessages {
		log.Println("send message:", msg)
		cSuccess <- subscribeServer.Send(msg)
	}
	return fmt.Errorf("subscriber closed")
}
