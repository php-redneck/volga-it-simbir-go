package controllers

import (
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/php-redneck/volga-it-simbir-go/internal/request"
	"github.com/php-redneck/volga-it-simbir-go/internal/responders"
	"github.com/php-redneck/volga-it-simbir-go/internal/services"
	"net/http"
	"strconv"
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
//	@Success		200
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
	log.Info(request.CreateTransportRequest(r))
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
func (c TransportController) Destroy(w http.ResponseWriter, r *http.Request) {

	responders.NoContent(w)
}
