package di

import (
	"database/sql"

	"github.com/rafaelsouzaribeiro/uow-golang/internal/infra/register"
	"github.com/rafaelsouzaribeiro/uow-golang/internal/usecase"
	"github.com/rafaelsouzaribeiro/uow-golang/pkg/uow"
)

func NewDI(con *sql.DB) *usecase.UseCase {
	uow := uow.NewUnitOfWork(con)
	register.RegisterRepositories(uow, con, "Repository")

	return usecase.NewUseCase(uow)
}
