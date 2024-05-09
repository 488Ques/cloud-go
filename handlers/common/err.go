package common

import (
	"net/http"
)

type ErrorCode string

type APIErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type statusError struct {
	error
	status int
}

func (e statusError) Unwrap() error {
	return e.error
}

func (e statusError) HTTPStatus() int {
	return e.status
}

var (
	ErrNotFound      = APIErrorResponse{Status: http.StatusNotFound, Message: "not found"}
	ErrDuplicate     = APIErrorResponse{Status: http.StatusBadRequest, Message: "duplicate"}
	ErrInternalError = APIErrorResponse{Status: http.StatusInternalServerError, Message: "internal server error"}
)

func ServerError(w http.ResponseWriter) {
	ServeJSON(w, http.StatusInternalServerError, ErrInternalError)
}
