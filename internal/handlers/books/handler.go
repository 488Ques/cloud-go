package books

import (
	"cloud-go/db"
	"cloud-go/internal/handlers/common"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	queries *db.Queries
	logger  *log.Logger
}

func New(queries *db.Queries, logger *log.Logger) *handler {
	return &handler{
		queries: queries,
		logger:  logger,
	}
}

func (h *handler) RegisterEndpoints(router *chi.Mux) {
	router.Route("/api/v1/book", func(r chi.Router) {
		r.Get("/", h.List)
	})
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	books, err := h.queries.ListPosts(r.Context())
	if err != nil {
		h.logger.Print(fmt.Errorf("ListPosts: failed getting a list of books: %w", err))
		common.ServerError(w)
		return
	}
	common.OKResponse(w, books)
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {}
func (h *handler) Read(w http.ResponseWriter, r *http.Request)   {}
func (h *handler) Update(w http.ResponseWriter, r *http.Request) {}
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {}
