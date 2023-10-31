package services

import (
	"github.com/php-redneck/volga-it-simbir-go/internal/entities"
	"github.com/php-redneck/volga-it-simbir-go/internal/repository"
)

type RentService struct {
	Repository repository.RentRepository
}

func (s RentService) AllByUser(id int) ([]entities.Rent, error) {
	return s.Repository.FindAll("user_id", id)
}

func (s RentService) AllByTransport(id int) ([]entities.Rent, error) {
	return s.Repository.FindAll("transport_id", id)
}

func (s RentService) Show(id int) (entities.Rent, error) {
	return s.Repository.FindBy("id", id)
}

func (s RentService) Destroy(id int) {

}
