package database

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/php-redneck/volga-it-simbir-go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instance *gorm.DB = nil

func Instance() *gorm.DB {
	if instance != nil {
		return instance
	}

	dbConfig := config.Database()

	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.DBName,
		dbConfig.Port,
		dbConfig.SSLMode,
		dbConfig.TimeZone,
	)

	if dbConfig.Password != "" {
		dsn += fmt.Sprintf(" password=%s", dbConfig.Password)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: nil,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	instance = db

	return instance
}
