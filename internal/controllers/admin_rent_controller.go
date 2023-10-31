package controllers

import (
	"github.com/php-redneck/volga-it-simbir-go/internal/services"
	"net/http"
)

type AdminRentController struct {
	Service services.UserService
}

// Index
// @Router			/api/Admin/Rent [get]
// @Tags			AdminRentController
// @Summary		Получение списка всех аккаунтов
// @Description	Получение списка всех аккаунтов
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Failure		401
// @Failure		500
func (c AdminRentController) Index(w http.ResponseWriter, r *http.Request) {}

// Show
// @Router			/api/Admin/Rent/{id} [get]
// @Tags			AdminRentController
// @Summary		Получение списка всех аккаунтов
// @Description	Получение списка всех аккаунтов
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Failure		401
// @Failure		500
func (c AdminRentController) Show(w http.ResponseWriter, r *http.Request) {}

// Store
// @Router			/api/Admin/Rent [post]
// @Tags			AdminRentController
// @Summary		Получение списка всех аккаунтов
// @Description	Получение списка всех аккаунтов
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Failure		401
// @Failure		500
func (c AdminRentController) Store(w http.ResponseWriter, r *http.Request) {}

// Edit
// @Router			/api/Admin/Rent/{id} [put]
// @Tags			AdminRentController
// @Summary		Получение списка всех аккаунтов
// @Description	Получение списка всех аккаунтов
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Failure		401
// @Failure		500
func (c AdminRentController) Edit(w http.ResponseWriter, r *http.Request) {}

// Destroy
// @Router			/api/Admin/Rent/{id} [delete]
// @Tags			AdminRentController
// @Summary		Получение списка всех аккаунтов
// @Description	Получение списка всех аккаунтов
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Failure		401
// @Failure		500
func (c AdminRentController) Destroy(w http.ResponseWriter, r *http.Request) {}
