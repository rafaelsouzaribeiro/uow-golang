package server

import (
	"github.com/rafaelsouzaribeiro/uow-golang/internal/infra/web/handler"
	"github.com/rafaelsouzaribeiro/uow-golang/internal/usecase"
)

func (s *Server) SetRoutes(usecase *usecase.UseCase) {
	handler := handler.NewHandler(usecase)
	s.chi.Post("/createBooks", handler.CreateAuthorAndBooks)
}
