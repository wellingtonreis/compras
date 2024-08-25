package main

import (
	"log"

	"compras/services/persist/config"
	"compras/services/persist/internal/routes"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	config.Load()
	app := fiber.New()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3003"))
}
