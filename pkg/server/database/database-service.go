package database

import (
	"fmt"
	"github.com/spike-events/spike-events/bin"
	"github.com/spike-events/spike-events/internal/env"
	"github.com/spike-events/spike-events/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

type Service interface {
	Subscribe(topicName, groupID string, offset int64) (models.Topic, error)
	TopicMessages(topic *models.Topic, groupID string, offset int64) ([]*bin.Message, error)
	CreateMessage(message *bin.Message) error
	UpdateTopics(topic, group string, offset int64) error
}

type srv struct {
	*gorm.DB
	m  sync.Mutex
	wg sync.WaitGroup
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
	return &srv{DB: db}
}

func defaultMigration(db *gorm.DB) error {
	return db.AutoMigrate(&models.Topic{})
}

func connectDatabase() (*gorm.DB, error) {
	connStr := os.Getenv("SPIKE_DB_CONN")
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
