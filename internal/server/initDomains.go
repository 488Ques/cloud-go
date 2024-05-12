package server

import (
	"cloud-go/internal/handlers/health"
	"cloud-go/internal/handlers/post"
)

func (s *Server) initDomains() {
	s.initHealthCheck()
	s.initPost()
}

func (s *Server) initHealthCheck() {
	healthCheckHandler := health.New()
	healthCheckHandler.RegisterEndpoints(s.router)
}

func (s *Server) initPost() {
	postHandler := post.New(s.queries, s.logger)
	postHandler.RegisterEndpoints(s.router)
}
