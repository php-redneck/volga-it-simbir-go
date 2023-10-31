package app

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	_ "github.com/php-redneck/volga-it-simbir-go/docs"
	"github.com/php-redneck/volga-it-simbir-go/internal/config"
	"github.com/php-redneck/volga-it-simbir-go/internal/entities"
	"github.com/php-redneck/volga-it-simbir-go/internal/middleware"
	"github.com/php-redneck/volga-it-simbir-go/internal/routes"
	"github.com/php-redneck/volga-it-simbir-go/pkg/database"
	"github.com/php-redneck/volga-it-simbir-go/pkg/httpserver"
	"github.com/swaggo/http-swagger"
)

func Run() {
	log.Info("open database connection")

	db := database.Instance()

	log.Info("start database auto migrations")

	err := db.AutoMigrate(
		entities.LogoutHistory{},
		entities.User{},
		entities.Transport{},
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	router := chi.NewRouter()

	router.Use(middleware.Global()...)

	router.Route("/api", routes.Api)

	router.Get(config.Swagger().BasePathPrefix+"/*", httpSwagger.Handler())

	server := httpserver.NewServer(config.App().PORT, router)

	server.Closer.Add(func(ctx context.Context) error {
		log.Info("close database connection")

		instance, err := db.DB()

		if err != nil {
			return err
		}

		return instance.Close()
	})

	server.Run()
}
