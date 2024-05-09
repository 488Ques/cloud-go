package common

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status int `json:"status"`
	Data   any `json:"data"`
}

func ServeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}

func OKResponse(w http.ResponseWriter, v any) {
	ServeJSON(w, http.StatusOK, APIResponse{
		Status: http.StatusOK,
		Data:   v,
	})
}
