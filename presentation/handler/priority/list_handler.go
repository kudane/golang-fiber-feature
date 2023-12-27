package handler

import (
	feature "backend/go-fiber/feature/priority"

	fiber "github.com/gofiber/fiber/v2"
)

func List(priority feature.Priority) fiber.Handler {
	return func(context *fiber.Ctx) error {
		result, _ := priority.List()
		return context.JSON(result)
	}
}
