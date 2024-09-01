package main

import (
	"log"

	"compras/services/persist/configs"
	"compras/services/persist/internal/routes"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	configs.LoadConfig()
	app := fiber.New()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":" + configs.Env.WebServerPort))
}
