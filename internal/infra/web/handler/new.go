package handler

import "github.com/rafaelsouzaribeiro/uow-golang/internal/usecase"

type Handler struct {
	usecase *usecase.UseCase
}

func NewHandler(usecase *usecase.UseCase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
