package handler

import (
	priority "backend/go-fiber/feature/priority"

	fiber "github.com/gofiber/fiber/v2"
)

func List(priority priority.List) fiber.Handler {
	return func(context *fiber.Ctx) error {
		result, _ := priority.List()
		return context.JSON(result)
	}
}
