package database

import (
	"log"
	"spike.io/internal/models"
)

func (s *srv) UpdateTopics(topic *models.Topic) error {
	err := s.Save(topic).Error
	if err != nil {
		log.Println(err)
	}
	return err
}
