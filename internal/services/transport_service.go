package services

import (
	"errors"
	"github.com/php-redneck/volga-it-simbir-go/internal/dto"
	"github.com/php-redneck/volga-it-simbir-go/internal/entities"
	"github.com/php-redneck/volga-it-simbir-go/internal/repository"
	"github.com/php-redneck/volga-it-simbir-go/pkg/database"
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

func (s TransportService) Index(paginationDto dto.TransportPaginationDTO) []entities.Transport {
	return s.Repository.Index(paginationDto.Start, paginationDto.Count, paginationDto.TransportType)
}

func (s TransportService) Show(id int) (entities.Transport, error) {
	return s.Repository.FindBy("id", id)
}

func (s TransportService) Store(dto dto.AdminTransportDTO) (entities.Transport, error) {
	var entity = entities.Transport{
		OwnerId:       dto.OwnerId,
		CanBeRented:   dto.CanBeRented,
		TransportType: dto.TransportType,
		Model:         dto.Model,
		Color:         dto.Color,
		Identifier:    dto.Identifier,
		Description:   dto.Description,
		Latitude:      dto.Latitude,
		Longitude:     dto.Longitude,
		MinutePrice:   dto.MinutePrice,
		DayPrice:      dto.DayPrice,
	}

	result := database.Instance().Create(&entity)

	if result.Error != nil {
		panic(result.Error)
	}

	return entity, nil
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
