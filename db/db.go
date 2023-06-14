package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBCredential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
}

func Connect(creds DBCredential) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,  // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,  // Don't include params in the SQL log
			Colorful:                  false, // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
