package routes

import (
	"compras/services/persist/internal/handlers"

	fiber "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/persist-data", handlers.Consumer)
}
