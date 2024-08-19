package routes

import (
	"compras/services/upload/internal/handlers"

	fiber "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/upload", handlers.Upload)
}
