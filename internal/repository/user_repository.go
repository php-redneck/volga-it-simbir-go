package repository

import (
	"fmt"
	"github.com/php-redneck/volga-it-simbir-go/internal/entities"
	"github.com/php-redneck/volga-it-simbir-go/pkg/database"
)

type UserRepository struct{}

func (r UserRepository) Index(start, count int) []entities.User {
	var users []entities.User

	database.Instance().Offset(start).Limit(count).Order("id desc").Find(&users)

	return users
}

func (r UserRepository) FindBy(field string, value interface{}) (entities.User, error) {
	user := entities.User{}

	res := database.Instance().First(&user, fmt.Sprintf("%s = ?", field), value)

	return user, res.Error
}
