package handler

import (
	feature "backend/go-fiber/feature/task"

	fiber "github.com/gofiber/fiber/v2"
)

func List(task feature.Task) fiber.Handler {
	return func(context *fiber.Ctx) error {
		result, _ := task.List()
		return context.JSON(result)
	}
}
