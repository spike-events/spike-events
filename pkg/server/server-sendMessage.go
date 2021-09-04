package server

import (
	"log"
	"spike.io/bin"
	"spike.io/internal/models"
	"sync"
)

func (s *server) sendMessage(message *bin.Message, newMessage bool) error {
	s.m.Lock()
	subs := s.subscribers[message.GetTopic()]
	s.m.Unlock()

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
			if message.Offset > sub.dbTopic.Offset[models.GroupID(message.GroupId)] {
				sub.dbTopic.Offset[models.GroupID(message.GroupId)] = message.Offset
			}
			wg.Done()
		}
	}()

	for _, item := range subs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := item.Send(message)
			if err != nil {
				// TODO: resend?
				return
			}
			wg.Add(1)
			cUpdateSent <- item
		}()
	}

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
