package handler

import (
	task "backend/go-fiber/core/task"

	fiber "github.com/gofiber/fiber/v2"
)

func ById(service task.Service) fiber.Handler {
	return func(context *fiber.Ctx) error {
		result, _ := service.ById(&task.CreateByIdCommand{Id: 1})
		return context.JSON(result)
	}
}
