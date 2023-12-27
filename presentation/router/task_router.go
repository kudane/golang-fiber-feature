package router

import (
	task "backend/go-fiber/core/task"
	handler "backend/go-fiber/presentation/handler/task"

	fiber "github.com/gofiber/fiber/v2"
)

func TaskRouter(app fiber.Router, service task.Service) {
	app.Get("/task/list", handler.List(service))
	app.Get("/task/:taskId", handler.ById(service))
}
