package controllers

import (
	"github.com/go-chi/chi/v5"
	"github.com/php-redneck/volga-it-simbir-go/internal/responders"
	"github.com/php-redneck/volga-it-simbir-go/internal/services"
	"net/http"
	"strconv"
)

type RentController struct {
	Service          services.RentService
	TransportService services.TransportService
}

// MyHistory
//
//	@Router			/api/Rent/MyHistory [get]
//	@Tags			RentController
//	@Summary		Получение истории аренд текущего аккаунта
//	@Description	Получение истории аренд текущего аккаунта
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Success		200 {object} []entities.Rent
//	@Failure		404
//	@Failure		500
func (c RentController) MyHistory(w http.ResponseWriter, r *http.Request) {
	list, err := c.Service.AllByUser(r.Context().Value("userId").(int))

	if err != nil {
		panic(err)
	}

	responders.OK(w, list)
}

// Show
//
//	@Router			/api/Rent/{id} [get]
//	@Tags			RentController
//	@Summary		Получение информации об аренде по id
//	@Description	Получение информации об аренде по id
//	@Security		BearerAuth
//	@Param			id	path	int	true	"Id аренды"
//	@Accept			json
//	@Produce		json
//	@Success		200 {object} entities.Rent
//	@Failure		404
//	@Failure		500
func (c RentController) Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		responders.NotFound(w)
		return
	}

	rent, err := c.Service.Show(id)

	if err != nil {
		responders.NotFound(w)
		return
	}

	userId := r.Context().Value("userId").(int)

	if rent.UserId != userId {
		transport, _ := c.TransportService.Show(rent.TransportId)

		if transport.OwnerId != userId {
			responders.NotFound(w)
			return
		}
	}

	responders.OK(w, rent)
}
