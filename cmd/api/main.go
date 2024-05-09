package main

import (
	"cloud-go/config"
	"cloud-go/db"
	"cloud-go/router"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type application struct {
	config  *config.Config
	logger  *log.Logger
	db      *pgx.Conn
	queries *db.Queries
	router  *chi.Mux
}

func main() {
	// Initialize the app's configuration
	config := config.New()

	// Initialize a new logger which writes messages to the standard output stream,
	// prefixed with the current date and time
	logger := log.New(os.Stdout, "", log.Ldate)

	// Initialize a connection to DB and construct a new queries instance
	// TODO Make this a connection pool
	conn, err := pgx.Connect(context.Background(), config.NewConnString())
	if err != nil {
		logger.Fatalf("pgx.Connect: unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())
	queries := db.New(conn)

	// Initialize router
	router := router.New(queries, logger)

	app := &application{
		config:  config,
		logger:  logger,
		db:      conn,
		queries: queries,
		router:  router,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
		Handler:      router,
		ReadTimeout:  config.Server.TimeoutRead,
		WriteTimeout: config.Server.TimeoutWrite,
		IdleTimeout:  config.Server.TimeoutIdle,
	}

	log.Printf("Starting server: %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
