package server

import (
	"log"
	"spike.io/bin"
	"spike.io/internal/models"
	"sync"
)

type balancerMessage struct {
}

func (s *server) sendMessage(message *bin.Message, newMessage bool) error {
	s.m.Lock()
	subs := s.subscribers[message.GetTopic()]
	subsNonPersistent := s.subscribersNonPersistence[message.GetTopic()]
	s.m.Unlock()

	var groupsSubs, groupsSubsNonPersistent map[string][]*subscribe

	for _, item := range subs {
		groupsSubs[item.topic.GroupId] = append(groupsSubs[item.topic.GroupId], item)
	}
	for _, item := range subsNonPersistent {
		groupsSubsNonPersistent[item.topic.GroupId] = append(groupsSubsNonPersistent[item.topic.GroupId], item)
	}

	if newMessage && len(subs) > 0 {
		err := s.db.CreateMessage(message)
		if err != nil {
			log.Println(err)
		}
	}

	var wg sync.WaitGroup
	cUpdateSent := make(chan *subscribe)
	go func() {
		for sub := range cUpdateSent {
			if message.Offset > *sub.dbTopic.Offset[models.GroupID(message.GroupId)] {
				offset := message.Offset
				sub.dbTopic.Offset[models.GroupID(message.GroupId)] = &offset
			}
			wg.Done()
		}
	}()

	next := func(subs []*subscribe) *subscribe {
		for _, n := range subs {
			if n.next {
				n.nextSub.next = true
				n.next = false
				return n
			}
		}
		return nil
	}
	send := func(subs map[string][]*subscribe, persistent bool) {
		for _, item := range subs {
			for i := 0; i < len(item); i++ {
				n := next(item)
				err := n.Send(message)
				if err != nil {
					// TODO: resend?
					continue
				}
				if persistent {
					wg.Add(1)
					cUpdateSent <- n
				}
				break
			}
		}
	}
	send(groupsSubs, true)
	send(groupsSubsNonPersistent, false)

	wg.Wait()
	close(cUpdateSent)

	if len(subs) > 0 {
		var topics []*models.Topic
		for _, t := range subs {
			topics = append(topics, t.dbTopic)
		}
		go s.db.UpdateTopics(topics)
	}

	return nil
}
