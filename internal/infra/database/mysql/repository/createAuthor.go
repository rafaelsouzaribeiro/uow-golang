package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/rafaelsouzaribeiro/uow-golang/internal/db"
	"github.com/rafaelsouzaribeiro/uow-golang/internal/entity"
)

func (r *Repository) CreateAuthor(argsAuthor entity.Author) (entity.Author, error) {
	ctx := context.Background()
	id := uuid.New().String()
	err := r.Queries.CreateAuthor(ctx, db.CreateAuthorParams{
		ID:   id,
		Name: argsAuthor.Name,
		Bio:  argsAuthor.Bio,
	})
	if err != nil {
		return entity.Author{}, err
	}
	return entity.Author{
		ID:   id,
		Name: argsAuthor.Name,
		Bio:  argsAuthor.Bio,
	}, nil
}
