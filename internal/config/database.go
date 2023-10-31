package config

import (
	"github.com/charmbracelet/log"
	"os"
	"strconv"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

func Database() DatabaseConfig {
	port := os.Getenv("DB_PORT")

	intPort, err := strconv.Atoi(port)

	if err != nil {
		log.Fatal("invalid database port")
	}

	return DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     intPort,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  "disable",
		TimeZone: "Europe/Moscow",
	}
}
