package routes

import (
	"comp/controllers"
	"comp/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRouter(app *fiber.App) {
	app.Post("/Product", middleware.JWTProtected(), controllers.CreateProduct)

	app.Post("/Category", middleware.JWTProtected(), controllers.CreateCategory)

	app.Get("/Product/:categoryId", controllers.GetAllProduct)

	app.Get("/Category", controllers.GetAllCategories)

	app.Get("/Product/:Id", controllers.GetProductById)
}