package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/php-redneck/volga-it-simbir-go/internal/app"
)

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	app.Run()
}
