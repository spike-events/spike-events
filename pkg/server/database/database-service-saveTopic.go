package database

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"spike.io/internal/models"
)

func (s *srv) UpdateTopics(topic, group string, offset int64) error {
	err := s.Model(&models.Topic{}).
		Where(&models.Topic{Name: topic}).
		Update("offset_data", gorm.Expr(fmt.Sprintf("jsonb_set(offset_data,'{%v}', '%v'::jsonb, true)", group, offset))).Error
	if err != nil {
		log.Println(err)
	}
	return err
}
