package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"os"
	"fmt"

	"comp/configs"
	"comp/routes"
)

func main() {
	uploadsDir := "./public"
	if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
		err := os.Mkdir(uploadsDir, 0755)
		if err != nil {
			fmt.Printf("Ошибка при создании папки: %v\n", err)
			return
		}
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Pivo"})
	})

	app.Static(
		"/public",  // mount address
		"./public", // path to the file folder
	)

	routes.FileRouter(app)

	routes.OrderController(app)

	routes.ProductRouter(app)

	routes.UserController(app)

	app.Listen(configs.EnvPort())
}