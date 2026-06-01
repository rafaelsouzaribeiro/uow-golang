package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/rafaelsouzaribeiro/uow-golang/internal/db"
	"github.com/rafaelsouzaribeiro/uow-golang/internal/entity"
)

func (r *Repository) CreateBook(argsBook entity.Books) (entity.Books, error) {
	ctx := context.Background()
	err := r.Queries.CreateBook(ctx, db.CreateBookParams{
		ID:          uuid.New().String(),
		Name:        argsBook.Name,
		Description: argsBook.Description,
		AuthorsID:   argsBook.AuthorsID,
		Price:       argsBook.Price,
	})
	if err != nil {
		return entity.Books{}, err
	}
	return entity.Books{
		Name:        argsBook.Name,
		Description: argsBook.Description,
		Price:       argsBook.Price,
		AuthorsID:   argsBook.AuthorsID,
	}, nil
}
