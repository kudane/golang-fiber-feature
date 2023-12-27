package handler

import (
	feature "backend/go-fiber/feature/task"

	fiber "github.com/gofiber/fiber/v2"
)

func ById(task feature.Task) fiber.Handler {
	return func(context *fiber.Ctx) error {
		result, _ := task.ById()
		return context.JSON(result)
	}
}
