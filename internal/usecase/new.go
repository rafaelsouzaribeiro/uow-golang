package usecase

import (
	"github.com/rafaelsouzaribeiro/uow-golang/pkg/uow"
)

type UseCase struct {
	uow uow.IunitOfWork
}

func NewUseCase(uow uow.IunitOfWork) *UseCase {
	return &UseCase{
		uow: uow,
	}
}
