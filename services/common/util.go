package util

import (
	"encoding/json"
	"net/http"
)

func ParseJSON(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteError(w http.ResponseWriter, status int, err error) {
	_ = WriteJSON(w, status, ErrorResponse{
		Error: err.Error(),
	})
}
