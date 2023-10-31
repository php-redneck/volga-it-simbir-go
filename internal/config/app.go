package config

import (
	"github.com/charmbracelet/log"
	"os"
	"strconv"
)

type AppConfig struct {
	PORT int
}

func App() *AppConfig {
	appPort := os.Getenv("APP_PORT")

	if appPort == "" {
		appPort = "3000"
	}

	port, err := strconv.Atoi(appPort)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return &AppConfig{
		PORT: port,
	}
}
