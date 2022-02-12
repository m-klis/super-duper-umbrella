package helpers

import (
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	StatusCode int         `json:"statusCode,omitempty"`
	Message    interface{} `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
	Count      interface{} `json:"count,omitempty"`
	Total      interface{} `json:"total,omitempty"`
}

// CustomResponse comment
func CustomResponse(w http.ResponseWriter, r *http.Request, code int, message interface{}, data interface{}) {
	render.Status(r, code)
	render.JSON(w, r, &Response{
		StatusCode: code,
		Message:    message,
		Data:       data,
	})
}

// Error Response comment
func ErrorResponse(w http.ResponseWriter, r *http.Request, code int, message string, err string) {
	render.Status(r, code)
	render.JSON(w, r, &Response{
		StatusCode: code,
		Message:    message,
		Error:      err,
	})
}
