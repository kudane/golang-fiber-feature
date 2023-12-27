package main

import (
	priority "backend/go-fiber/feature/priority"
	task "backend/go-fiber/feature/task"
	database "backend/go-fiber/infrastructure/database"
	smtp "backend/go-fiber/infrastructure/smtp"
	router "backend/go-fiber/presentation/router"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// infra
	db := database.SqlLite{}
	smtp := smtp.Smtp{}

	// feature: task
	task := task.New(&db, &smtp)
	router.TaskRouter(app, task)

	// feature: priority
	priority := priority.New(&db)
	router.PriorityRouter(app, priority)

	app.Listen(":3000")
}
