package entity

import (
	"database/sql"

	"github.com/google/uuid"
)

type Author struct {
	ID   string
	Name string
	Bio  sql.NullString
}

func NewAuthor(name string, bio sql.NullString) Author {
	return Author{
		ID:   uuid.New().String(),
		Name: name,
		Bio:  bio,
	}
}
