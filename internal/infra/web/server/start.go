package server

import "net/http"

func (s *Server) Start(port string) error {
	return http.ListenAndServe(":"+port, s.chi)
}
