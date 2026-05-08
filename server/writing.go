package server

import (
	"encoding/json"
	"io"
	"net/http"
)

func Decode(r io.Reader, data any) error {
	return json.NewDecoder(r).Decode(data)
}

func WriteError(w http.ResponseWriter, err string, statusCode int) {
	errDTO := NewErrorDTO(err)
	http.Error(w, errDTO.ToString(), statusCode)
}

func WriteJSON(w http.ResponseWriter, data any, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
