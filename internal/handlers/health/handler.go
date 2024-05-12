package health

import (
	"cloud-go/internal/handlers/common"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (h *handler) RegisterEndpoints(router *chi.Mux) {
	router.Route("/api/v1/healthcheck", func(r chi.Router) {
		r.Get("/", h.Read)
	})
}

func (h *handler) Read(w http.ResponseWriter, r *http.Request) {
	common.OKResponse(w, "ok")
}
