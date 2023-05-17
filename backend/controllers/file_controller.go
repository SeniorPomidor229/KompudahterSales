package controllers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func UploudImage(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Ошибка при получении файла")
	}

	err = c.SaveFile(file, fmt.Sprintf("./public/%s", file.Filename))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Ошибка при сохранении файла")
	}

	filePath := fmt.Sprintf("/public/%s", file.Filename)

	return c.Status(http.StatusOK).JSON(&fiber.Map{"data": filePath})
}
