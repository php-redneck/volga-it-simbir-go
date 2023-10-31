package controllers

import (
	"github.com/go-chi/chi/v5"
	"github.com/php-redneck/volga-it-simbir-go/internal/dto"
	"github.com/php-redneck/volga-it-simbir-go/internal/responders"
	"github.com/php-redneck/volga-it-simbir-go/internal/services"
	"net/http"
	"strconv"
)

type PaymentController struct {
	UserService services.UserService
}

// Hesoyam
//
//	@Router			/api/Payment/Hesoyam/{accountId} [post]
//	@Tags			PaymentController
//	@Summary		Добавляет на баланс аккаунта с id={accountId} 250 000 денежных единиц.
//	@Description	Администратор может добавить баланс всем, обычный пользователь только себе
//	@Param			accountId	path	int	true	"ID пользователя, если отправляет не администратор, то поле игонируется"
//	@Security		BearerAuth
//	@Success		204
//	@Failure		401
//	@Failure		500
func (c PaymentController) Hesoyam(w http.ResponseWriter, r *http.Request) {
	user, err := c.UserService.Show(r.Context().Value("userId").(int))

	if err != nil {
		panic(err)
	}

	if !user.IsAdmin {
		_, err = c.UserService.Edit(int(user.ID), dto.EditUserDTO{
			Username: "",
			Password: "",
			IsAdmin:  user.IsAdmin,
			Balance:  user.Balance + 250000,
		})

		if err != nil {
			panic(err)
		}

		responders.NoContent(w)

		return
	}

	accountId, err := strconv.Atoi(chi.URLParam(r, "accountId"))

	if err != nil {
		responders.NotFound(w)
		return
	}

	user, err = c.UserService.Show(accountId)

	if err != nil {
		responders.NotFound(w)
		return
	}

	_, err = c.UserService.Edit(int(user.ID), dto.EditUserDTO{
		Username: "",
		Password: "",
		IsAdmin:  user.IsAdmin,
		Balance:  user.Balance + 250000,
	})

	if err != nil {
		panic(err)
	}

	responders.NoContent(w)
}
