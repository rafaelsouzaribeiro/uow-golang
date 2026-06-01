package irepository

import "github.com/rafaelsouzaribeiro/uow-golang/internal/entity"

type IRepository interface {
	CreateAuthor(argsAuthor entity.Author) (entity.Author, error)
	CreateBook(argsBook entity.Books) (entity.Books, error)
}
