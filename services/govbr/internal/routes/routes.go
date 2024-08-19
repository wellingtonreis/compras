package routes

import (
	"compras/services/govbr/internal/handlers"

	fiber "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Consumer)
}
