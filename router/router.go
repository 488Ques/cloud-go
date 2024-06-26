package router

import (
	"cloud-go/resources/book"
	"cloud-go/resources/health"

	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthcheck", health.Read)

	r.Route("/v1", func(r chi.Router) {
		bookAPI := &book.API{}
		r.Get("/books", bookAPI.List)
		r.Post("/books", bookAPI.Create)
		r.Get("/books/{id}", bookAPI.Read)
		r.Put("/books/{id}", bookAPI.Update)
		r.Delete("/books/{id}", bookAPI.Delete)
	})

	return r
}
