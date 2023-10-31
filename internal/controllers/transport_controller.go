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
// @Router			/api/Transport/{id} [get]
// @Tags			TransportController
// @Summary		Получение информации о транспорте по id
// @Description	Получение информации о транспорте по id
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
//	@Summary		Получение информации о транспорте по id
//	@Description	Получение информации о транспорте по id
func (c TransportController) Store(w http.ResponseWriter, r *http.Request) {
	log.Info(request.CreateTransportRequest(r))
}

// Edit
// @Router			/api/Transport/{id} [put]
// @Tags			TransportController
// @Summary		Получение информации о транспорте по id
// @Description	Получение информации о транспорте по id
func (c TransportController) Edit(w http.ResponseWriter, r *http.Request) {

}

// Destroy
// @Router			/api/Transport/{id} [delete]
// @Tags			TransportController
// @Summary		Получение информации о транспорте по id
// @Description	Получение информации о транспорте по id
func (c TransportController) Destroy(w http.ResponseWriter, r *http.Request) {

	responders.NoContent(w)
}
