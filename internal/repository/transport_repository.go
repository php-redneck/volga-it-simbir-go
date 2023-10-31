package repository

import (
	"fmt"
	"github.com/php-redneck/volga-it-simbir-go/internal/entities"
	"github.com/php-redneck/volga-it-simbir-go/pkg/database"
)

type TransportRepository struct{}

func (r TransportRepository) Index(start, count int, transportType string) []entities.Transport {
	var users []entities.Transport

	query := database.Instance().Offset(start).Limit(count).Order("id desc")

	if transportType != "All" {
		query.Where("transport_type = ?", transportType)
	}

	query.Find(&users)

	return users
}

func (r TransportRepository) FindBy(field string, value interface{}) (entities.Transport, error) {
	user := entities.Transport{}

	res := database.Instance().First(&user, fmt.Sprintf("%s = ?", field), value)

	return user, res.Error
}
