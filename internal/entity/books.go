package entity

import (
	"database/sql"

	"github.com/google/uuid"
)

type Books struct {
	ID          string
	Name        string
	Description sql.NullString
	AuthorsID   string
	Price       float64
}

func NewBooks(name string, description sql.NullString, authorsID string, price float64) Books {
	return Books{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		AuthorsID:   authorsID,
		Price:       price,
	}
}
