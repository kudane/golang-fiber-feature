package feature

import (
	database "backend/go-fiber/infrastructure/database"
)

type Priority interface {
	List
}

type priority struct {
	db *database.SqlLite
}

func New(db *database.SqlLite) Priority {
	return &priority{
		db: db,
	}
}
