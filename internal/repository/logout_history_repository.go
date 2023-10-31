package repository

import (
	"fmt"
	"github.com/php-redneck/volga-it-simbir-go/internal/entities"
	"github.com/php-redneck/volga-it-simbir-go/pkg/database"
)

type LogoutHistoryRepository struct{}

func (r LogoutHistoryRepository) FindBy(key string, value interface{}) (entities.LogoutHistory, error) {
	var entity entities.LogoutHistory

	res := database.Instance().First(&entity, fmt.Sprintf("%s = ?", key), value)

	return entity, res.Error
}
