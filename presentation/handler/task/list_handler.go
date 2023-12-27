package handler

import (
	task "backend/go-fiber/feature/task"

	fiber "github.com/gofiber/fiber/v2"
)

func List(task task.List) fiber.Handler {
	return func(context *fiber.Ctx) error {
		result, _ := task.List()
		return context.JSON(result)
	}
}
