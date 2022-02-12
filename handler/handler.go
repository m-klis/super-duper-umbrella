package handler

import (
	"gochicoba/helpers"
	"net/http"
)

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	helpers.ErrorResponse(w, r, http.StatusMethodNotAllowed, "failed", "nil")
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	helpers.ErrorResponse(w, r, http.StatusNotFound, "failed", "nil")
}
