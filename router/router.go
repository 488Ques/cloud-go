package router

import (
	"cloud-go/db"
	"cloud-go/handlers/book"
	"cloud-go/handlers/health"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(queries *db.Queries, logger *log.Logger) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/healthcheck", health.Read)

	bookAPI := book.NewHandler(queries, logger)

	bookRouter := chi.NewRouter()
	bookRouter.Get("/", bookAPI.List)
	bookRouter.Post("/", bookAPI.Create)
	bookRouter.Get("/{id}", bookAPI.Read)
	bookRouter.Put("/{id}", bookAPI.Update)
	bookRouter.Delete("/{id}", bookAPI.Delete)

	apiRouter := chi.NewRouter()
	apiRouter.Mount("/books", bookRouter)
	r.Mount("/v1", apiRouter)

	return r
}
