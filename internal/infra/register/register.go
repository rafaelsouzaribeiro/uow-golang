package register

import (
	"database/sql"

	"github.com/rafaelsouzaribeiro/uow-golang/internal/db"
	"github.com/rafaelsouzaribeiro/uow-golang/internal/infra/database/mysql/repository"
	"github.com/rafaelsouzaribeiro/uow-golang/pkg/uow"
)

func RegisterRepositories(unit *uow.UnitOfWork, dbConn *sql.DB, name string) {
	unit.Register(name, func(tx *sql.Tx) interface{} {
		repo := repository.NewRepository(dbConn)
		repo.Queries = db.New(tx)
		return repo
	})
}
