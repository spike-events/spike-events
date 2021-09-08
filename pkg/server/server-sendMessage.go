package server

import (
	"fmt"
	"github.com/google/uuid"
	"spike.io/bin"
	"spike.io/internal/models"
	"sync"
)

type balancerMessage struct {
}

func (s *server) sendMessage(message *bin.Message, newMessage bool, id uuid.UUID) error {
	s.m.Lock()
	subs := s.subscribers[message.GetTopic()]
	subsNonPersistent := s.subscribersNonPersistence[message.GetTopic()]
	s.m.Unlock()

	groupsSubs, groupsSubsNonPersistent := make(map[string][]*subscribe), make(map[string][]*subscribe)

	for _, item := range subs {
		if !newMessage && item.id != id {
			continue
		}
		groupsSubs[item.topic.GroupId] = append(groupsSubs[item.topic.GroupId], item)
	}
	for _, item := range subsNonPersistent {
		if !newMessage && item.id != id {
			continue
		}
		groupsSubsNonPersistent[item.topic.GroupId] = append(groupsSubsNonPersistent[item.topic.GroupId], item)
	}

	if newMessage && len(subs) > 0 {
		err := s.db.CreateMessage(subs[0].dbTopic, message)
		if err != nil {
			fmt.Println(err)
		}
	}

	var wg sync.WaitGroup
	cUpdateSent := make(chan *subscribe)
	go func() {
		for sub := range cUpdateSent {
			offset := sub.dbTopic.Offset[models.GroupID(sub.topic.GetGroupId())]
			if offset == nil || message.Offset > *offset {
				offset := message.Offset
				sub.dbTopic.Offset[models.GroupID(sub.topic.GetGroupId())] = &offset
			}
			wg.Done()
		}
	}()

	next := func(subs []*subscribe) *subscribe {
		for _, n := range subs {
			if n.next {
				n.nextSub.next = true
				if len(subs) > 1 {
					n.next = false
				}
				return n
			}
		}
		return nil
	}
	send := func(subs map[string][]*subscribe, persistent bool) {
		for _, item := range subs {
			for i := 0; i < len(item); i++ {
				n := next(item)
				fmt.Println("send message channel:", message)
				n.msg <- message
				fmt.Println("read message channel:", message)
				err := <-n.success
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
