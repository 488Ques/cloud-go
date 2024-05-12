package server

import (
	"cloud-go/internal/handlers/books"
	"cloud-go/internal/handlers/health"
)

func (s *Server) initDomains() {
	s.initHealthCheck()
	s.initBooks()
}

func (s *Server) initHealthCheck() {
	healthCheckHandler := health.New()
	healthCheckHandler.RegisterEndpoints(s.router)
}

func (s *Server) initBooks() {
	booksHandler := books.New(s.queries, s.logger)
	booksHandler.RegisterEndpoints(s.router)
}
