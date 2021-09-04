package database

import (
	"spike.io/bin"
	"spike.io/internal/models"
)

func (s *srv) CreateMessage(message *bin.Message) error {
	msg := models.Message{
		Message: *message,
	}
	err := s.Create(&msg).Error
	if err != nil {
		return err
	}
	message.Offset = msg.ID
	return nil
}
