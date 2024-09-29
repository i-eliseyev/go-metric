package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/i-eliseyev/go-metric/internal/handlers"
)

func SetupRouters(app *fiber.App) {
	app.Get("/", handlers.HandleIndex)

	update := app.Group("/update")
	update.Post("/", handlers.HandleUpdateMetric)

	value := app.Group("/value")
	value.Get("/", handlers.HandleGetMetric)
}
