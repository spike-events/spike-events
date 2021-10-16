package server

import (
	"fmt"
	"github.com/spike-events/spike-events/bin"
	"log"
)

type balancerMessage struct {
}

func (s *server) sendMessage(message *bin.Message, newMessage bool, id string) error {
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

	if newMessage {
		err := s.db.CreateMessage(message)
		if err != nil {
			fmt.Println(err)
		}
	}

	next := func(subs []*subscribe) *subscribe {
		for i, n := range subs {
			if n.next {
				n.next = false
				if len(subs)-1 == i {
					subs[0].next = true
				} else {
					subs[i+1].next = true
				}
				return n
			}
		}
		return nil
	}
	send := func(subs map[string][]*subscribe, persistent bool) error {
		log.Println(">>> server send: ", message.GetTopic(), message)
		var err error
		for _, item := range subs {
			for i := 0; i < len(item); i++ {
				n := next(item)
				if n == nil {
					fmt.Println(">>> server topic not fount: ", message.GetOffset(), message.GetTopic())
					continue
				}

				fmt.Println(">>> server send message channel:", message.GetOffset(), message.GetTopic())
				n.msg <- message
				fmt.Println("<<< server send message channel:", message.GetOffset(), message.GetTopic())

				fmt.Println(">>> server wait ok:", message.GetOffset(), message.GetTopic())
				err = <-n.success
				fmt.Println("<<< server wait ok:", message.GetOffset(), message.GetTopic())

				if err != nil {
					// TODO: resend?
					continue
				}

				err = nil
				if persistent {
					s.db.UpdateTopics(message.GetTopic(), n.topic.GroupId, message.GetOffset())
				}
				break
			}
		}
		return err
	}
	err := send(groupsSubs, true)
	send(groupsSubsNonPersistent, false)

	if err != nil {
		return err
	}

	return nil
}
