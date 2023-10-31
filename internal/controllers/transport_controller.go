package controllers

import (
	"github.com/go-chi/chi/v5"
	"github.com/php-redneck/volga-it-simbir-go/internal/dto"
	"github.com/php-redneck/volga-it-simbir-go/internal/request"
	"github.com/php-redneck/volga-it-simbir-go/internal/responders"
	"github.com/php-redneck/volga-it-simbir-go/internal/services"
	"net/http"
	"strconv"
	"strings"
)

type TransportController struct {
	Service services.TransportService
}

// Show
//
//	@Router			/api/Transport/{id} [get]
//	@Tags			TransportController
//	@Summary		Получение информации о транспорте по id
//	@Description	Получение информации о транспорте по id
//
//	@Param			id	path	int	true	"Id транспорта"
//	@Accept			json
//	@Produce		json
//	@Success		200 {object} entities.Transport
//	@Failure		404
//	@Failure		500
func (c TransportController) Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		responders.NotFound(w)
		return
	}

	entity, err := c.Service.Show(id)

	if err != nil {
		responders.NotFound(w)
		return
	}

	responders.OK(w, entity)
}

// Store
//
//	@Router			/api/Transport [post]
//	@Tags			TransportController
//	@Summary		Добавление нового транспорта
//	@Description	Добавление нового транспорта
//	@Param			body	body	dto.TransportDTO	true	"Параметры запроса"
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
func (c TransportController) Store(w http.ResponseWriter, r *http.Request) {
	tdto, err := request.CreateTransportRequest(r)

	if err != nil {
		responders.BadRequest(w, strings.Split(err.Error(), "\n"))
		return
	}

	admDto := dto.AdminTransportDTO{
		OwnerId:       r.Context().Value("userId").(int),
		CanBeRented:   tdto.CanBeRented,
		TransportType: tdto.TransportType,
		Model:         tdto.Model,
		Color:         tdto.Color,
		Identifier:    tdto.Identifier,
		Description:   tdto.Description,
		Latitude:      tdto.Latitude,
		Longitude:     tdto.Longitude,
		MinutePrice:   tdto.MinutePrice,
		DayPrice:      tdto.DayPrice,
	}

	entity, err := c.Service.Store(admDto)

	if err != nil {
		panic(err)
	}

	responders.OK(w, entity)
}

// Edit
//
//	@Router			/api/Transport/{id} [put]
//	@Tags			TransportController
//	@Summary		Изменение транспорта оп id
//	@Description	Изменение транспорта оп id
func (c TransportController) Edit(w http.ResponseWriter, r *http.Request) {

}

// Destroy
//
//	@Router			/api/Transport/{id} [delete]
//	@Tags			TransportController
//	@Summary		Удаление транспорта по id
//	@Description	Удаление транспорта по id
//	@Param			id	path	int	true	"Id транспорта"
//	@Accept			json
//	@Produce		json
//	@Success		204
//	@Failure		403
//	@Failure		404
//	@Failure		500
func (c TransportController) Destroy(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		responders.NotFound(w)
		return
	}

	entity, err := c.Service.Show(id)

	if err != nil {
		responders.NotFound(w)
		return
	}

	if entity.OwnerId != r.Context().Value("userId").(int) {
		responders.Forbidden(w)
	}

	err = c.Service.Destroy(id)

	if err != nil {
		panic(err)
	}

	responders.NoContent(w)
}
