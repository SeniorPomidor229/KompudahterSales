package routes

import (
	"comp/controllers"

	"github.com/gofiber/fiber/v2"
)

func FileRouter(app *fiber.App) {
	app.Post("/Uploud", controllers.UploudImage)
}