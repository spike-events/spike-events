package database

import (
	"log"
	"spike.io/internal/models"
)

func (s *srv) UpdateTopics(topics []*models.Topic) {
	err := s.Save(&topics).Error
	if err != nil {
		log.Println(err)
	}
}
