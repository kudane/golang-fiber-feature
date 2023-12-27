package handler

import (
	task "backend/go-fiber/feature/task"

	fiber "github.com/gofiber/fiber/v2"
)

func ById(task task.ById) fiber.Handler {
	return func(context *fiber.Ctx) error {
		result, _ := task.ById()
		return context.JSON(result)
	}
}
