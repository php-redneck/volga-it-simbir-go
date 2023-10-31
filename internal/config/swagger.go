package config

import (
	"github.com/php-redneck/volga-it-simbir-go/docs"
	"github.com/swaggo/swag"
)

type SwaggerConfig struct {
	Info           *swag.Spec
	BasePathPrefix string
}

func Swagger() SwaggerConfig {
	docs.SwaggerInfo.Title = "Simbir.GO"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "API спецификация сервиса для аренды различных транспортных средств"

	return SwaggerConfig{
		Info:           docs.SwaggerInfo,
		BasePathPrefix: "/api/docs",
	}
}
