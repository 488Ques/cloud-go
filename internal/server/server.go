package server

import (
	"cloud-go/db"
	"cloud-go/internal/config"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type Server struct {
	config     *config.Config
	logger     *log.Logger
	router     *chi.Mux
	conn       *pgx.Conn
	queries    *db.Queries
	httpServer *http.Server
}

func New() *Server {
	s := &Server{
		config: config.New(),
		logger: log.New(os.Stdout, "", log.Ldate),
		router: chi.NewRouter(),
	}
	s.init()
	return s
}

func (s *Server) Run() {
	s.httpServer = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", s.config.Server.Host, s.config.Server.Port),
		Handler:      s.router,
		ReadTimeout:  s.config.Server.TimeoutRead,
		WriteTimeout: s.config.Server.TimeoutWrite,
		IdleTimeout:  s.config.Server.TimeoutIdle,
	}
	log.Printf("Starting server: %s", s.httpServer.Addr)
	if err := s.httpServer.ListenAndServe(); err != nil {
		s.logger.Fatalf("Server.Run: unable to start server: %v\n", err)
	}
	defer s.cleanup()
}

func (s *Server) init() {
	s.initDatabase()
	s.initDomains()
}

func (s *Server) initDatabase() {
	// TODO Make this a connection pool
	conn, err := pgx.Connect(context.Background(), s.config.NewConnString())
	if err != nil {
		s.logger.Fatalf("Server.initDatabase: unable to connect to database: %v\n", err)
	}
	s.conn = conn
	s.queries = db.New(conn)
}

func (s *Server) cleanup() {
	if err := s.conn.Close(context.Background()); err != nil {
		s.logger.Fatalf("Server.cleanup: cannot close pgx.Conn: %v\n", err)
	}
}
