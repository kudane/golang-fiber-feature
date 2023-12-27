package feature

import (
	database "backend/go-fiber/infrastructure/database"
	smtp "backend/go-fiber/infrastructure/smtp"
)

type Task interface {
	List
	ById
}

type task struct {
	db   *database.SqlLite
	smtp *smtp.Smtp
}

func New(db *database.SqlLite, smtp *smtp.Smtp) Task {
	return &task{
		db:   db,
		smtp: smtp,
	}
}
