package database

import (
	"fmt"
	"github.com/spike-events/spike-events/internal/models"
	"gorm.io/gorm"
	"log"
)

func (s *srv) UpdateTopics(topic, group string, offset int64) error {
	defer s.m.Unlock()
	s.m.Lock()

	err := s.Model(&models.Topic{}).
		Where(&models.Topic{Name: topic}).
		Update("offset_data", gorm.Expr(fmt.Sprintf("jsonb_set(offset_data,'{%v}', '%v'::jsonb, true)", group, offset))).Error
	if err != nil {
		log.Println(err)
	}
	return err
}
