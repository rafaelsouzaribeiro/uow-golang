package server

import "github.com/go-chi/chi"

type Server struct {
	chi *chi.Mux
}

func NewServer() *Server {
	chiRouter := chi.NewRouter()
	return &Server{
		chi: chiRouter,
	}
}
