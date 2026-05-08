package server

import (
	"encoding/json"
	"net/http"
	"rest-api/books"
)

type HTTPHandler struct {
	library *books.Library
}

func NewHTTPHandler(library *books.Library) *HTTPHandler {
	return &HTTPHandler{
		library: library,
	}
}

func (h *HTTPHandler) AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var bookDTO BookDTO
	if err := json.NewDecoder(r.Body).Decode(&bookDTO); err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := bookDTO.Validate(); err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	book := books.NewBook(bookDTO.Title, bookDTO.Author, bookDTO.Year)
	if err := h.library.AddBook(book); err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := WriteJSON(w, book); err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}
}
