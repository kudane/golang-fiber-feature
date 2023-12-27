package handler

import (
	task "backend/go-fiber/core/task"

	fiber "github.com/gofiber/fiber/v2"
)

func List(service task.Service) fiber.Handler {
	return func(context *fiber.Ctx) error {
		result, _ := service.List(&task.CreateListCommand{})
		return context.JSON(result)
	}
}
