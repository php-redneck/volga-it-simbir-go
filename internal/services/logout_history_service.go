package services

import (
	"github.com/php-redneck/volga-it-simbir-go/internal/entities"
	"github.com/php-redneck/volga-it-simbir-go/internal/repository"
	"github.com/php-redneck/volga-it-simbir-go/pkg/database"
)

type LogoutHistoryService struct {
	Repository repository.LogoutHistoryRepository
}

// IsLogout - Check user jwt token id on exists in logout history table
func (s LogoutHistoryService) IsLogout(jti string) bool {
	_, err := s.Repository.FindBy("jti", jti)
	return err == nil
}

// Create - Add jwt token id to logout history table
func (s LogoutHistoryService) Create(jti string) {
	entity := entities.LogoutHistory{JTI: jti}

	if res := database.Instance().Create(&entity); res.Error != nil {
		panic(res.Error.Error())
	}
}
