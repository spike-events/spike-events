package database

import (
	"fmt"
	"github.com/google/uuid"
	"spike.io/bin"
	"spike.io/internal/models"
	"strings"
)

var TemplateTableName = "topic_messages_%v"

func (s *srv) Subscribe(topicName, groupID string, offset int64) (models.Topic, error) {
	var topic models.Topic

	err := s.Where(&models.Topic{Name: topicName}).First(&topic).Error
	if err != nil {
		topic.Offset = make(map[models.GroupID]*int64)
		//new subscribe
		id := strings.ReplaceAll(uuid.New().String(), "-", "")

		var zero int64 = 0
		topic.Name = topicName
		topic.Offset[models.GroupID(groupID)] = &zero
		topic.Table = fmt.Sprintf(TemplateTableName, id)
		err = s.Create(&topic).Error
		if err != nil {
			return topic, err
		}

		err = s.Table(topic.Table).AutoMigrate(&models.Message{})
		if err != nil {
			return topic, err
		}

		return topic, nil
	}

	return topic, err
}

func (s *srv) TopicMessages(topic models.Topic, groupID string, offset int64) ([]*bin.Message, error) {
	var result []*bin.Message

	// load messages
	if offset > 0 {
		var messages []models.Message
		s.Table(topic.Table).Where("id >= ?", offset).Order("id").Find(&messages)
		for _, item := range messages {
			result = append(result, &item.Message)
		}
	} else {
		lastOffset := topic.Offset[models.GroupID(groupID)]
		var messages []models.Message
		s.Table(topic.Table).Where("id >= ?", lastOffset).Order("id").First(&messages)
		for _, item := range messages {
			result = append(result, &item.Message)
		}
	}

	// update group
	exists := topic.Offset[models.GroupID(groupID)]
	if exists == nil {
		var zero int64 = 0
		topic.Offset[models.GroupID(groupID)] = &zero
		err := s.Save(&topic).Error
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
