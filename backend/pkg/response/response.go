package response

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// JSON writes a JSON response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

// Error writes an error response
func Error(w http.ResponseWriter, statusCode int, err error) error {
	return JSON(w, statusCode, ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: err.Error(),
	})
}

// Success writes a successful response
func Success(w http.ResponseWriter, data interface{}) error {
	return JSON(w, http.StatusOK, data)
}
