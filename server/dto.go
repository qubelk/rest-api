package server

import (
	"encoding/json"
	"rest-api/books"
	"rest-api/sins"
	"time"
)

type BookDTO struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func BookFromDTO(b BookDTO) books.Book {
	return books.NewBook(b.Title, b.Author, b.Year)
}

func (b BookDTO) Validate() error {
	switch {
	case b.Title == "":
		return sins.InvalidTitleArgument
	case b.Author == "":
		return sins.InvalidAuthorArgument
	case b.Year > time.Now().Year():
		return sins.InvalidYearArgument
	}

	return nil
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func NewErrorDTO(message string) ErrorDTO {
	return ErrorDTO{
		Message: message,
		Time:    time.Now(),
	}
}

func (e ErrorDTO) ToString() string {
	data, _ := json.MarshalIndent(e, "", "    ")
	return string(data)
}
