package router

import (
	feature "backend/go-fiber/feature/task"
	handler "backend/go-fiber/presentation/handler/task"

	fiber "github.com/gofiber/fiber/v2"
)

func TaskRouter(app fiber.Router, task feature.Task) {
	app.Get("/task/list", handler.List(task))
	app.Get("/task/:taskId", handler.ById(task))
}
