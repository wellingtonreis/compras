package main

import (
	"compras/services/govbr/config"
	"compras/services/govbr/internal/routes"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	config.Load()
	app := fiber.New()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3002"))
}
