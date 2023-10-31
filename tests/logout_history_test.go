package tests

import (
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/php-redneck/volga-it-simbir-go/internal/services"
	"github.com/php-redneck/volga-it-simbir-go/pkg/database"
	"testing"
)

func TestCreateLogoutHistory(t *testing.T) {
	godotenv.Load("../.env")

	db := database.Instance()

	defer func() {
		instance, err := db.DB()

		if err != nil {
			panic(err)
		}

		if err = instance.Close(); err != nil {
			return
		}
	}()

	db.Begin()

	id := uuid.New().String()

	service := services.LogoutHistoryService{}

	service.Create(id)

	if !service.IsLogout(id) {
		log.Fatal("jwt id does not exists in logout table")
	}

	db.Rollback()
}
