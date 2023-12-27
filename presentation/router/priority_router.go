package router

import (
	priority "backend/go-fiber/core/priority"
	handler "backend/go-fiber/presentation/handler/priority"

	fiber "github.com/gofiber/fiber/v2"
)

func PriorityRouter(app fiber.Router, service priority.Service) {
	app.Get("/priority/list", handler.List(service))
}
