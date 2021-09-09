package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"spike.io/bin"
	"spike.io/internal/env"
	"spike.io/internal/models"
)

type Service interface {
	Subscribe(topicName, groupID string, offset int64) (models.Topic, error)
	TopicMessages(topic models.Topic, groupID string, offset int64) ([]*bin.Message, error)
	CreateMessage(message *bin.Message) error
	UpdateTopics(topic *models.Topic) error
}

type srv struct {
	*gorm.DB
}

func New() Service {
	db, err := connectDatabase()
	if err != nil {
		panic(err)
	}
	err = defaultMigration(db)
	if err != nil {
		panic(err)
	}
	return &srv{db}
}

func defaultMigration(db *gorm.DB) error {
	return db.AutoMigrate(&models.Topic{})
}

func connectDatabase() (*gorm.DB, error) {
	connStr := os.Getenv("DB_CONN")
	if len(connStr) == 0 {
		return nil, fmt.Errorf("no DB_CONN environment variable provided")
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  env.ConnectionString,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	base, err := db.DB()
	if err != nil {
		return nil, err
	}
	base.SetMaxIdleConns(env.MaxConnections)
	base.SetMaxOpenConns(env.MaxConnections)
	return db, nil
}
