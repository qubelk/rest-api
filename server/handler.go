package server

import (
	"net/http"
	"rest-api/books"

	"github.com/go-chi/chi/v5"
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
	if err := Decode(r.Body, &bookDTO); err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := bookDTO.Validate(); err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	book := BookFromDTO(bookDTO)
	if err := h.library.AddBook(&book); err != nil {
		WriteError(w, err.Error(), http.StatusConflict)
		return
	}

	if err := WriteJSON(w, book, http.StatusCreated); err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *HTTPHandler) GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books := h.library.GetAllBooks()
	if err := WriteJSON(w, books, http.StatusOK); err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HTTPHandler) GetAllAvailableBooksHandler(w http.ResponseWriter, r *http.Request) {
	books := h.library.GetAllAvailableBooks()
	if err := WriteJSON(w, books, http.StatusOK); err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HTTPHandler) GetAllBorrowedBooksHandler(w http.ResponseWriter, r *http.Request) {
	books := h.library.GetAllBorrowedBooks()
	if err := WriteJSON(w, books, http.StatusOK); err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HTTPHandler) GetBookByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data, err := h.library.GetBookByID(id)
	if err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	WriteJSON(w, data, http.StatusOK)
}

func (h *HTTPHandler) BorrowBookHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.library.BorrowBook(id); err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *HTTPHandler) ReturnBookHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.library.ReturnBook(id); err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *HTTPHandler) DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.library.DeleteBook(id); err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
