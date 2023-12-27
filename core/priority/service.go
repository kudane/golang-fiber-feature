package core

import (
	database "backend/go-fiber/infrastructure/database"
)

type Service interface {
	List(*CreateListCommand) (*CreateListCommandResponse, error)
}

type service struct {
	db *database.SqlLite
}

func New(db *database.SqlLite) Service {
	return &service{
		db: db,
	}
}
