package database

import (
	"github.com/spike-events/spike-events/bin"
	"github.com/spike-events/spike-events/internal/models"
)

func (s *srv) CreateMessage(message *bin.Message) error {
	defer s.m.Unlock()
	s.m.Lock()

	var dbTopic models.Topic
	err := s.Where(&models.Topic{Name: message.GetTopic()}).First(&dbTopic).Error
	if err == nil {
		msg := models.Message{
			Message: message,
		}
		err := s.Debug().Table(dbTopic.Table).Create(&msg).Error
		if err != nil {
			return err
		}
		message.Offset = msg.ID
	}
	return nil
}
