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

type AdminAccountController struct {
	Service services.UserService
}

// Index
//
//	@Router			/api/Admin/Account [get]
//	@Tags			AdminAccountController
//	@Summary		Получение списка всех аккаунтов
//	@Description	Получение списка всех аккаунтов
//	@Param			start	query	int	true	"Начало выборки"
//	@Param			count	query	int	true	"Размер выборки"
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	responders.PaginationResponse{payload=[]entities.User}
//	@Failure		400	{object}	responders.BadRequestResponse
//	@Failure		401
//	@Failure		403
//	@Failure		500
func (c AdminAccountController) Index(w http.ResponseWriter, r *http.Request) {
	paginationDTO, err := request.PaginationRequest(r)

	if err != nil {
		responders.BadRequest(w, strings.Split(err.Error(), "\n"))
		return
	}

	total := c.Service.Total()

	list := c.Service.Index(paginationDTO.Start, paginationDTO.Count)

	var data = make([]interface{}, len(list))

	for i, user := range list {
		data[i] = user
	}

	responders.Pagination(w, data, responders.PaginationShortMeta{
		Total:        int(total),
		Count:        paginationDTO.Count,
		Start:        paginationDTO.Start,
		CountRecords: len(list),
	})
}

// Show
//
//	@Router			/api/Admin/Account/{id} [get]
//	@Tags			AdminAccountController
//	@Summary		Получение информации об аккаунте по id
//	@Description	Получение информации об аккаунте по id
//	@Param			id	path	int	true	"Id пользователя"
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	entities.User
//	@Failure		401
//	@Failure		403
//	@Failure		404
//	@Failure		500
func (c AdminAccountController) Show(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		responders.NotFound(w)
		return
	}

	user, err := c.Service.Show(userId)

	if err != nil {
		responders.NotFound(w)
		return
	}

	responders.OK(w, user)
}

// Store
//
//	@Router			/api/Admin/Account [post]
//	@Tags			AdminAccountController
//	@Summary		Создание администратором нового аккаунта
//	@Description	Создание администратором нового аккаунта
//	@Param			body	body	dto.CreateUserDTO	true	"Параметры запроса"
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		201	{object}	entities.User
//	@Failure		400	{object}	responders.BadRequestResponse
//	@Failure		401
//	@Failure		403
//	@Failure		500
func (c AdminAccountController) Store(w http.ResponseWriter, r *http.Request) {
	dto, err := request.CreateUserRequest(r)

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
//	@Router			/api/Admin/Account/{id} [put]
//	@Tags			AdminAccountController
//	@Summary		Изменение администратором аккаунта по id
//	@Description	Изменение администратором аккаунта по id
//	@Param			id		path	int				true	"Id пользователя"
//	@Param			body	body	dto.EditUserDTO	true	"Параметры запроса"
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		204
//	@Failure		400	{object}	responders.BadRequestResponse
//	@Failure		401
//	@Failure		403
//	@Failure		404
//	@Failure		500
func (c AdminAccountController) Edit(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		responders.NotFound(w)
		return
	}

	editUserDTO, err := request.EditUserRequest(r)

	if err != nil {
		responders.BadRequest(w, strings.Split(err.Error(), "\n"))
		return
	}

	if _, err = c.Service.Edit(userId, editUserDTO); err != nil {
		if string(err.Error()) == "not_found" {
			responders.NotFound(w)
		} else {
			responders.BadRequest(w, strings.Split(err.Error(), "\n"))
		}
		return
	}

	responders.NoContent(w)
}

// Destroy
//
//	@Router			/api/Admin/Account/{id} [delete]
//	@Tags			AdminAccountController
//	@Summary		Удаление аккаунта по id
//	@Description	Удаление аккаунта по id
//	@Param			id	path	int	true	"Id пользователя"
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		204
//	@Failure		401
//	@Failure		403
//	@Failure		404
//	@Failure		500
func (c AdminAccountController) Destroy(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		responders.NotFound(w)
		return
	}

	if err := c.Service.Destroy(userId); err != nil {
		responders.NotFound(w)
		return
	}

	responders.NoContent(w)
}
