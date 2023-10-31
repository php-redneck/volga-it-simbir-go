package controllers

import (
	"github.com/go-chi/chi/v5"
	"github.com/php-redneck/volga-it-simbir-go/internal/request"
	"github.com/php-redneck/volga-it-simbir-go/internal/responders"
	"github.com/php-redneck/volga-it-simbir-go/internal/services"
	"net/http"
	"strconv"
	"strings"
)

type AdminTransportController struct {
	Service services.TransportService
}

// Index
//
//	@Router			/api/Admin/Transport [get]
//	@Tags			AdminTransportController
//	@Summary		Получение списка всех транспортных средств
//	@Description	Получение списка всех транспортных средств
//
//	@Param			start			query	int		true	"Начало выборки"
//	@Param			count			query	int		true	"Размер выборки"
//	@Param			transportType	query	string	true	"Тип транспорта [Car, Bike, Scooter, All]"
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Failure		400	{object}	responders.BadRequestResponse
//	@Failure		401
//	@Failure		403
//	@Failure		500
func (c AdminTransportController) Index(w http.ResponseWriter, r *http.Request) {
	transportPaginationDTO, err := request.PaginationAdminTransportRequest(r)

	if err != nil {
		responders.BadRequest(w, strings.Split(err.Error(), "\n"))
		return
	}

	total := c.Service.Total(transportPaginationDTO.TransportType)

	list := c.Service.Index(transportPaginationDTO)

	var data = make([]interface{}, len(list))

	for i, user := range list {
		data[i] = user
	}

	responders.Pagination(w, data, responders.PaginationShortMeta{
		Total:        int(total),
		Count:        transportPaginationDTO.Count,
		Start:        transportPaginationDTO.Start,
		CountRecords: len(list),
	})
}

// Show
//
//	@Router			/api/Admin/Transport/{id} [get]
//	@Tags			AdminTransportController
//	@Summary		Получение информации о транспортном средстве по id
//	@Description	Получение информации о транспортном средстве по id
//	@Param			id	path	int	true	"Id транспорта"
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Failure		401
//	@Failure		403
//	@Failure		404
//	@Failure		500
func (c AdminTransportController) Show(w http.ResponseWriter, r *http.Request) {
	transportId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		responders.NotFound(w)
		return
	}

	transport, err := c.Service.Show(transportId)

	if err != nil {
		responders.NotFound(w)
		return
	}

	responders.OK(w, transport)
}

// Store
//
//	@Router			/api/Admin/Transport [post]
//	@Tags			AdminTransportController
//	@Summary		Создание нового транспортного средства
//	@Description	Создание нового транспортного средства
//	@Param			body	body	dto.AdminTransportDTO	true	"Параметры запроса"
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Failure		400	{object}	responders.BadRequestResponse
//	@Failure		401
//	@Failure		403
//	@Failure		500
func (c AdminTransportController) Store(w http.ResponseWriter, r *http.Request) {
	dto, err := request.AdminCreateTransportRequest(r)

	if err != nil {
		responders.BadRequest(w, strings.Split(err.Error(), "\n"))
		return
	}

	entity, err := c.Service.Store(dto)

	if err != nil {
		responders.BadRequest(w, []string{err.Error()})
		return
	}

	responders.OK(w, entity)
}

// Edit
//
//	@Router			/api/Admin/Transport/{id} [put]
//	@Tags			AdminTransportController
//	@Summary		Изменение транспортного средства по id
//	@Description	Изменение транспортного средства по id
//	@Param			id		path	int						true	"Id пользователя"
//	@Param			body	body	dto.AdminTransportDTO	true	"Параметры запроса"
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		204
//	@Failure		400	{object}	responders.BadRequestResponse
//	@Failure		401
//	@Failure		403
//	@Failure		404
//	@Failure		500
func (c AdminTransportController) Edit(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		responders.NotFound(w)
		return
	}

	editTransportDTO, err := request.EditTransportRequest(r)

	if err != nil {
		responders.BadRequest(w, strings.Split(err.Error(), "\n"))
		return
	}

	if _, err = c.Service.Edit(id, editTransportDTO); err != nil {
		panic(err)
	}

	responders.NoContent(w)
}

// Destroy
//
//	@Router			/api/Admin/Transport/{id} [delete]
//	@Tags			AdminTransportController
//	@Summary		Удаление транспорта по id
//	@Description	Удаление транспорта по id
//	@Param			id	path	int	true	"Id транспорта"
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		204
//	@Failure		401
//	@Failure		403
//	@Failure		404
//	@Failure		500
func (c AdminTransportController) Destroy(w http.ResponseWriter, r *http.Request) {
	transportId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		responders.NotFound(w)
		return
	}

	if err := c.Service.Destroy(transportId); err != nil {
		responders.NotFound(w)
		return
	}

	responders.NoContent(w)
}
