package database

import (
	"spike.io/bin"
	"spike.io/internal/models"
)

func (s *srv) CreateMessage(topic *models.Topic, message *bin.Message) error {
	msg := models.Message{
		Message: *message,
	}
	err := s.Table(topic.Table).Create(&msg).Error
	if err != nil {
		return err
	}
	message.Offset = msg.ID
	return nil
}
