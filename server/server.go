package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	handlers *HTTPHandler
}

func NewServer(handlers *HTTPHandler) *Server {
	return &Server{
		handlers: handlers,
	}
}

func (s *Server) StartServer() error {
	r := chi.NewRouter()

	r.Route("/books", func(r chi.Router) {
		r.Post("/", s.handlers.AddBookHandler)
		r.Get("/", s.handlers.GetAllBooksHandler)
		r.Get("/available", s.handlers.GetAllAvailableBooksHandler)
		r.Get("/borrowed", s.handlers.GetAllBorrowedBooksHandler)

		r.Get("/{id}", s.handlers.GetBookByIDHandler)
		r.Patch("/{id}/borrows", s.handlers.BorrowBookHandler)
		r.Patch("/{id}/return", s.handlers.ReturnBookHandler)
		r.Delete("/{id}", s.handlers.DeleteBookHandler)
	})

	log.Println("Server started on localhost:3000/books")
	if err := http.ListenAndServe(":3000", r); err != nil {
		return err
	}
	return nil
}
