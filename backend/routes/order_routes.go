package routes

import (
	"comp/controllers"
	"comp/middleware"

	"github.com/gofiber/fiber/v2"
)

func OrderController(app *fiber.App) {
	app.Post("/Order", middleware.JWTProtected(), controllers.CreateOrder)

	app.Get("/Order", middleware.JWTProtected(), controllers.GetOrders)
}