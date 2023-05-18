package routes

import (
	"comp/controllers"
	//"comp/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserController(app *fiber.App) {
	app.Post("/Login", controllers.Login)

	app.Post("/Register", controllers.Register)
}