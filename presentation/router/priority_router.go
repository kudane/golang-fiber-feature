package router

import (
	feature "backend/go-fiber/feature/priority"
	handler "backend/go-fiber/presentation/handler/priority"

	fiber "github.com/gofiber/fiber/v2"
)

func PriorityRouter(app fiber.Router, priority feature.Priority) {
	app.Get("/priority/list", handler.List(priority))
}
