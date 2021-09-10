package test

import (
	"fmt"
	"github.com/spike-events/spike-events/pkg/client"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

// ConnectDatabase connect dabase postgres
func connectDatabase() (*gorm.DB, error) {
	env := os.Getenv("DB_CONN")
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  env,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		// TODO: Foreign keys requires migration dependencies which are not supported yet
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   newLogger,
	})
	if err != nil {
		return nil, err
	}
	base, err := db.DB()
	if err != nil {
		return nil, err
	}
	maxConns, err := strconv.Atoi(os.Getenv("DB_MAX_CONNS"))
	if err != nil || maxConns == 0 {
		maxConns = 3
	}
	base.SetMaxIdleConns(maxConns)
	base.SetMaxOpenConns(maxConns)
	return db, nil
}

func cleanDatabase() {
	db, err := connectDatabase()

	if err != nil {
		panic("Unable to clear database")
	}

	type tableNames struct {
		TableName string
	}
	var tables []tableNames
	db.Raw(`SELECT table_name 
		FROM information_schema.tables
		WHERE table_schema='public'
		AND table_type='BASE TABLE'
		ORDER BY table_schema,table_name`).Scan(&tables)

	for _, table := range tables {
		err = db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", table.TableName)).Error
		if err != nil {
			panic(err)
		}
	}
}

var spikeConn client.Conn

func TestMain(m *testing.M) {
	cleanDatabase()

	result := m.Run()
	os.Exit(result)
}
