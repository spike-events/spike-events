package server

import (
	"fmt"
	"github.com/spike-events/spike-events/bin"
	"github.com/spike-events/spike-events/internal/models"
	"log"
)

type subscribe struct {
	id      string
	msg     chan *bin.Message
	success chan error
	topic   *bin.Topic
	next    bool
}

func (s *server) registerSubscribe(topic *bin.Topic) (chan *bin.Message, chan error) {
	defer s.m.Unlock()
	s.m.Lock()
	if s.subscribers == nil {
		s.subscribers = make(map[string][]*subscribe)
	}
	if s.subscribersNonPersistence == nil {
		s.subscribersNonPersistence = make(map[string][]*subscribe)
	}
	isNext := func(subs []*subscribe) bool {
		if len(subs) == 0 {
			return true
		}
		var countGroup int
		for _, item := range subs {
			if item.topic.GetGroupId() == topic.GetGroupId() {
				countGroup++
			}
		}
		if countGroup == 0 {
			return true
		}
		return false
	}
	channel := make(chan *bin.Message, 100)
	success := make(chan error, 100)
	if topic.Persistent {
		last := &subscribe{topic.GetId(), channel, success, topic, isNext(s.subscribers[topic.Topic])}
		s.subscribers[topic.Topic] = append(s.subscribers[topic.Topic], last)
		go func() {
			dbTopic := s.registerSubscribeDatabase(topic)
			messages, err := s.db.TopicMessages(&dbTopic, topic.GroupId, topic.Offset)
			if err != nil {
				log.Println(err)
			}
			for _, item := range messages {
				s.sendMessage(item, false, topic.GetId())
			}
		}()
	} else {
		last := &subscribe{topic.GetId(), channel, success, topic, isNext(s.subscribersNonPersistence[topic.Topic])}
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
