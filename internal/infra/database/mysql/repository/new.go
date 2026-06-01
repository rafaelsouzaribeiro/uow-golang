package repository

import (
	"database/sql"

	"github.com/rafaelsouzaribeiro/uow-golang/internal/db"
)

type Repository struct {
	dbConn  *sql.DB
	Queries *db.Queries
}

func NewRepository(dbConn *sql.DB) *Repository {
	return &Repository{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}
