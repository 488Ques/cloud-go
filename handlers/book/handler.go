package book

import (
	"cloud-go/db"
	"cloud-go/handlers/common"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	queries *db.Queries
	logger  *log.Logger
}

func NewHandler(queries *db.Queries, logger *log.Logger) *Handler {
	return &Handler{
		queries: queries,
		logger:  logger,
	}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	books, err := h.queries.ListPosts(r.Context())
	if err != nil {
		h.logger.Print(fmt.Errorf("ListPosts: failed getting a list of books: %w", err))
		common.ServerError(w)
		return
	}
	common.OKResponse(w, books)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Read(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {}
