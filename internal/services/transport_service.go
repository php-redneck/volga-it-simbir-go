package services

import (
	"errors"
	"github.com/php-redneck/volga-it-simbir-go/internal/dto"
	"github.com/php-redneck/volga-it-simbir-go/internal/entities"
	"github.com/php-redneck/volga-it-simbir-go/internal/repository"
	"github.com/php-redneck/volga-it-simbir-go/pkg/database"
	"gorm.io/gorm"
)

type TransportService struct {
	Repository repository.TransportRepository
}

func (s TransportService) Total(transportType string) int64 {
	var (
		count     int64
		transport entities.Transport
		query     = database.Instance().Model(&transport)
	)

	if transportType != "all" {
		query.Where("transport_type = ?", transportType)
	}

	query.Count(&count)

	return count
}

func (s TransportService) FindByUsername(username string) (entities.Transport, error) {
	return s.Repository.FindBy("username", username)
}

func (s TransportService) Index(paginationDto dto.TransportPaginationDTO) []entities.Transport {
	return s.Repository.Index(paginationDto.Start, paginationDto.Count, paginationDto.TransportType)
}

func (s TransportService) Show(id int) (entities.Transport, error) {
	return s.Repository.FindBy("id", id)
}

func (s TransportService) Store(dto dto.AdminTransportDTO) (entities.Transport, error) {
	entity := map[string]interface{}{}

	if dto.Description == "" {
		entity["description"] = gorm.Expr("NULL")
	}

	if dto.MinutePrice == 0 {
		entity["minute_price"] = gorm.Expr("NULL")
	}

	if dto.DayPrice == 0 {
		entity["day_price"] = gorm.Expr("NULL")
	}

	var model = entities.Transport{
		OwnerId:       dto.OwnerId,
		CanBeRented:   dto.CanBeRented,
		TransportType: dto.TransportType,
		Model:         dto.Model,
		Color:         dto.Color,
		Identifier:    dto.Identifier,
		//Description:   dto.Description,
		Latitude:  dto.Latitude,
		Longitude: dto.Longitude,
		//MinutePrice:   dto.MinutePrice,
		//DayPrice:      dto.DayPrice,
	}

	result := database.Instance().Attrs(entity).Create(&model)

	if result.Error != nil {
		panic(result.Error)
	}

	return model, nil
}

func (s TransportService) Edit(id int, dto dto.TransportDTO) (entities.Transport, error) {
	transport, err := s.Show(id)

	if err != nil {
		return transport, errors.New("not_found")
	}

	transport.TransportType = dto.TransportType

	transport.Identifier = dto.Identifier

	transport.Color = dto.Color

	transport.Model = dto.Model

	transport.DayPrice = dto.DayPrice

	transport.MinutePrice = dto.MinutePrice

	transport.Description = dto.Description

	transport.Longitude = dto.Longitude

	transport.Latitude = dto.Latitude

	transport.CanBeRented = dto.CanBeRented

	res := database.Instance().Save(&transport)

	return transport, res.Error
}

func (s TransportService) Destroy(id int) error {
	user, err := s.Show(id)

	if err != nil {
		return err
	}

	database.Instance().Delete(&user)

	return nil
}
