package uow

import (
	"context"
	"database/sql"
	"fmt"
)

type RepositoryTx func(tx *sql.Tx) interface{}

type IunitOfWork interface {
	Register(name string, fn RepositoryTx) interface{}
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Commit() error
	Rollback() error
	UnRegister(name string)
	Do(ctx context.Context, fu func(uow *UnitOfWork) error) error
}

type UnitOfWork struct {
	Db           *sql.DB
	Repositories map[string]RepositoryTx
	Tx           *sql.Tx
}

func NewUnitOfWork(db *sql.DB) *UnitOfWork {
	return &UnitOfWork{
		Db:           db,
		Repositories: make(map[string]RepositoryTx),
	}
}
func (u *UnitOfWork) Register(name string, fn RepositoryTx) interface{} {
	u.Repositories[name] = fn
	return fn
}

func (u *UnitOfWork) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if u.Tx == nil {
		tx, err := u.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		u.Tx = tx
	}

	repoFactory, exists := u.Repositories[name]
	if !exists {
		return nil, fmt.Errorf("repository %s not registered", name)
	}

	return repoFactory(u.Tx), nil
}

func (u *UnitOfWork) Commit() error {
	if u.Tx != nil {
		return u.Tx.Commit()
	}
	return nil
}

func (u *UnitOfWork) Rollback() error {
	if u.Tx != nil {
		return u.Tx.Rollback()
	}
	return nil
}

func (u *UnitOfWork) UnRegister(name string) {
	delete(u.Repositories, name)
}

func (u *UnitOfWork) Do(ctx context.Context, fu func(uow *UnitOfWork) error) error {
	err := fu(u)
	if err != nil {
		u.Rollback()
		return err
	}
	return u.Commit()
}
