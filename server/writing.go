package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Decode(r io.Reader, data any) error {
	return json.NewDecoder(r).Decode(data)
}

func WriteError(w http.ResponseWriter, err string, statusCode int) {
	errDTO := NewErrorDTO(err)
	http.Error(w, errDTO.ToString(), statusCode)
}

func WriteJSON(w http.ResponseWriter, data any) error {
	b, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		log.Printf("Failed to write http response, error: %s\n", err.Error())
		return nil
	}
	return nil
}
