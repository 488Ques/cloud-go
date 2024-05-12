package books

import (
	"cloud-go/db"
	"cloud-go/internal/handlers/common"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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
		r.Post("/", h.Create)
	})
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	posts, err := h.queries.ListPosts(r.Context())
	if err != nil {
		h.logger.Printf("queries.ListPosts: failed to get a list of posts: %v", err)
		common.ServerError(w)
		return
	}
	common.OKResponse(w, posts)
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var postRequest struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	err := json.NewDecoder(r.Body).Decode(&postRequest)
	if err != nil {
		return
	}

	uuid := uuid.New()
	post, err := h.queries.CreatePost(r.Context(), db.CreatePostParams{
		ID:      uuid,
		Title:   postRequest.Title,
		Content: postRequest.Content,
	})
	if err != nil {
		h.logger.Printf("queries.CreatePost: failed to create a post: %v", err)
		common.ServerError(w)
		return
	}
	common.OKResponse(w, post)
}

func (h *handler) Read(w http.ResponseWriter, r *http.Request)   {}
func (h *handler) Update(w http.ResponseWriter, r *http.Request) {}
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {}
