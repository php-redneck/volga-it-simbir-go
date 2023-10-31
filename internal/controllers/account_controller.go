package controllers

import (
	"github.com/php-redneck/volga-it-simbir-go/internal/request"
	"github.com/php-redneck/volga-it-simbir-go/internal/responders"
	"github.com/php-redneck/volga-it-simbir-go/internal/services"
	"net/http"
	"strings"
)

type AccountController struct {
	Service              services.AccountService
	LogoutHistoryService services.LogoutHistoryService
}

// Me
//
//	@Router			/api/Account/Me [get]
//	@Tags			AccountController
//	@Summary		Получение данных о текущем аккаунте
//	@Description	Получение данных о текущем аккаунте
//	@Security		BearerAuth
//	@Success		200	{object}	entities.User
//	@Failure		401
//	@Failure		500
func (c AccountController) Me(w http.ResponseWriter, r *http.Request) {
	user, err := c.Service.UserService.Show(r.Context().Value("userId").(int))

	if err != nil {
		panic(err)
	}

	responders.OK(w, user)
}

// SignIn
//
//	@Router			/api/Account/SignIn [post]
//	@Tags			AccountController
//	@Summary		Получение нового jwt токена пользователя
//	@Description	Авторизация пользователя
//	@Param			body	body	dto.SignInDTO	true	"Параметры запроса"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string	"Токен"
//	@Failure		400	{object}	responders.BadRequestResponse
//	@Failure		401
//	@Failure		500
func (c AccountController) SignIn(w http.ResponseWriter, r *http.Request) {
	signInDTO, err := request.SignInRequest(r)

	if err != nil {
		responders.BadRequest(w, strings.Split(err.Error(), "\n"))
		return
	}

	token, err := c.Service.SignIn(signInDTO)

	if err != nil {
		responders.BadRequest(w, strings.Split(err.Error(), "\n"))
		return
	}

	responders.OK(w, token)
}

// SignUp
//
//	@Router			/api/Account/SignUp [post]
//	@Tags			AccountController
//	@Summary		Регистрация нового аккаунта
//	@Description	Регистрация пользователя
//	@Param			body	body	dto.SignUpDTO	true	"Параметры запроса"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string	"Токен"
//	@Failure		400	{object}	responders.BadRequestResponse
//	@Failure		401
//	@Failure		500
func (c AccountController) SignUp(w http.ResponseWriter, r *http.Request) {
	signUpDTO, err := request.SignUpRequest(r)

	if err != nil {
		responders.BadRequest(w, strings.Split(err.Error(), "\n"))
		return
	}

	token, err := c.Service.SignUp(signUpDTO)

	if err != nil {
		responders.BadRequest(w, strings.Split(err.Error(), "\n"))
		return
	}

	responders.OK(w, token)
}

// SignOut
//
//	@Router			/api/Account/SignOut [post]
//	@Tags			AccountController
//	@Summary		Выход из аккаунта
//	@Description	Выход из аккаунта
//	@Security		BearerAuth
//	@Success		204
//	@Failure		401
//	@Failure		500
func (c AccountController) SignOut(w http.ResponseWriter, r *http.Request) {
	c.LogoutHistoryService.Create(r.Context().Value("tokenId").(string))

	responders.NoContent(w)
}

// Update
//
//	@Router			/api/Account/Update [put]
//	@Tags			AccountController
//	@Summary		Обновление своего аккаунта
//	@Description	Обновление своего аккаунта
//	@Param			body	body	dto.AccountUpdateDTO	true	"Параметры запроса"
//	@Security		BearerAuth
//	@Success		204
//	@Failure		401
//	@Failure		500
func (c AccountController) Update(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(int)

	accountUpdateDTO, err := request.AccountUpdateRequest(r)

	if err != nil {
		responders.BadRequest(w, strings.Split(err.Error(), "\n"))
		return
	}

	_, err = c.Service.Update(userId, accountUpdateDTO)

	if err != nil {
		responders.BadRequest(w, strings.Split(err.Error(), "\n"))
		return
	}

	responders.NoContent(w)
}
