package core

import (
	database "backend/go-fiber/infrastructure/database"
	smtp "backend/go-fiber/infrastructure/smtp"
)

type Service interface {
	List(*CreateListCommand) (*CreateListCommandResponse, error)
	ById(*CreateByIdCommand) (*CreateByIdCommandResponse, error)
}

type service struct {
	db   *database.SqlLite
	smtp *smtp.Smtp
}

func New(db *database.SqlLite, smtp *smtp.Smtp) Service {
	return &service{
		db:   db,
		smtp: smtp,
	}
}
