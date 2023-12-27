package handler

import (
	priority "backend/go-fiber/core/priority"

	fiber "github.com/gofiber/fiber/v2"
)

func List(service priority.Service) fiber.Handler {
	return func(context *fiber.Ctx) error {
		result, _ := service.List(&priority.CreateListCommand{})
		return context.JSON(result)
	}
}
