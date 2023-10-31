package repository

import (
	"fmt"
	"github.com/php-redneck/volga-it-simbir-go/internal/entities"
	"github.com/php-redneck/volga-it-simbir-go/pkg/database"
)

type RentRepository struct{}

func (r RentRepository) FindAll(field string, value interface{}) ([]entities.Rent, error) {
	var rents []entities.Rent

	res := database.Instance().Order("id desc").Find(&rents, fmt.Sprintf("%s = ?", field), value)

	return rents, res.Error
}

func (r RentRepository) FindBy(field string, value interface{}) (entities.Rent, error) {
	rent := entities.Rent{}

	res := database.Instance().First(&rent, fmt.Sprintf("%s = ?", field), value)

	return rent, res.Error
}
