package database

import (
	"spike.io/bin"
	"spike.io/internal/models"
)

func (s *srv) CreateMessage(message *bin.Message) error {
	var dbTopic models.Topic
	err := s.Where(&models.Topic{Name: message.GetTopic()}).First(&dbTopic).Error
	if err == nil {
		msg := models.Message{
			Message: message,
		}
		err := s.Table(dbTopic.Table).Create(&msg).Error
		if err != nil {
			return err
		}
		message.Offset = msg.ID
	}
	return nil
}
