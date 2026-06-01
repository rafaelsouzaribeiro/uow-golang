package usecase

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rafaelsouzaribeiro/uow-golang/internal/dto"
	"github.com/rafaelsouzaribeiro/uow-golang/internal/entity"
	"github.com/rafaelsouzaribeiro/uow-golang/internal/irepository"
	"github.com/rafaelsouzaribeiro/uow-golang/pkg/uow"
)

func (u *UseCase) CreateAuthorAndBook(args dto.AuthorBooks) error {
	ctx := context.Background()
	if args.AuthorName == "" {
		return fmt.Errorf("author Name is required")
	}
	if args.BookName == "" || args.BookPrice == 0 {
		return fmt.Errorf("book Name and Price are required")
	}

	err := u.uow.Do(ctx, func(uow *uow.UnitOfWork) error {
		var err error

		repo, err := uow.GetRepository(ctx, "Repository")
		if err != nil {
			return err
		}

		author := entity.Author{
			Name: args.AuthorName,
			Bio:  sql.NullString{String: args.AuthorBio, Valid: true},
		}

		at, err := repo.(irepository.IRepository).CreateAuthor(author)

		if err != nil {
			return err
		}

		book := entity.Books{
			Name:      args.BookName,
			Price:     args.BookPrice,
			AuthorsID: at.ID,
		}

		_, err = repo.(irepository.IRepository).CreateBook(book)
		if err != nil {
			return err
		}
		return nil

	})
	if err != nil {
		return err
	}
	return nil
}
