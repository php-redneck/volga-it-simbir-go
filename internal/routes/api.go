package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/php-redneck/volga-it-simbir-go/internal/controllers"
	"github.com/php-redneck/volga-it-simbir-go/internal/middleware"
)

func Api(r chi.Router) {

	r.Route("/Account", func(r chi.Router) {

		r.Group(func(r chi.Router) {
			r.Use(middleware.ValidateJWT)

			r.Get("/Me", controllers.AccountController{}.Me)

			r.Post("/SignOut", controllers.AccountController{}.SignOut)

			r.Put("/Update", controllers.AccountController{}.Update)
		})

		r.Post("/SignIn", controllers.AccountController{}.SignIn)

		r.Post("/SignUp", controllers.AccountController{}.SignUp)

	})

	r.Route("/Transport", func(r chi.Router) {

		r.Get("/{id}", controllers.TransportController{}.Show)

		r.Group(func(r chi.Router) {
			r.Use(middleware.ValidateJWT)

			r.Post("/", controllers.TransportController{}.Store)

			r.Put("/{id}", controllers.TransportController{}.Edit)

			r.Delete("/{id}", controllers.TransportController{}.Destroy)
		})

	})

	r.Route("/Rent", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middleware.ValidateJWT)

			r.Get("/MyHistory", controllers.RentController{}.MyHistory)

			r.Get("/{id}", controllers.RentController{}.Show)
		})
	})

	r.Route("/Payment", func(r chi.Router) {
		r.Use(middleware.ValidateJWT)

		r.Post("/Hesoyam/{accountId}", controllers.PaymentController{}.Hesoyam)
	})

	r.Route("/Admin", func(r chi.Router) {
		r.Use(middleware.ValidateJWT, middleware.IsAdmin)

		r.Route("/Account", func(r chi.Router) {

			r.Get("/", controllers.AdminAccountController{}.Index)

			r.Post("/", controllers.AdminAccountController{}.Store)

			r.Get("/{id}", controllers.AdminAccountController{}.Show)

			r.Put("/{id}", controllers.AdminAccountController{}.Edit)

			r.Delete("/{id}", controllers.AdminAccountController{}.Destroy)

		})

		r.Route("/Transport", func(r chi.Router) {

			r.Get("/", controllers.AdminTransportController{}.Index)

			r.Post("/", controllers.AdminTransportController{}.Store)

			r.Get("/{id}", controllers.AdminTransportController{}.Show)

			r.Put("/{id}", controllers.AdminTransportController{}.Edit)

			r.Delete("/{id}", controllers.AdminTransportController{}.Destroy)

		})

		//r.Route("/Rent", func(r chi.Router) {
		//
		//	r.Get("/", controllers.AdminTransportController{}.Index)
		//
		//	r.Post("/", controllers.AdminTransportController{}.Store)
		//
		//	r.Get("/{id}", controllers.AdminTransportController{}.Show)
		//
		//	r.Put("/{id}", controllers.AdminTransportController{}.Edit)
		//
		//	r.Delete("/{id}", controllers.AdminTransportController{}.Destroy)
		//
		//})
	})

}
