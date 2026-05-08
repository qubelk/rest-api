package books

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Author      string     `json:"author"`
	Year        int        `json:"year"`
	IsAvailable bool       `json:"isAvailable"`
	AddedAt     time.Time  `json:"addedAt"`
	BorrowedAt  *time.Time `json:"borrowedAt"`
}

func validate(title string, author string, year int) error {
	switch {
	case title == "":
		return InvalidTitleArgument
	case author == "":
		return InvalidAuthorArgument
	case year > time.Now().Year():
		return InvalidYearArgument
	}

	return nil
}

func NewBook(title string, author string, year int) (Book, error) {
	if err := validate(title, author, year); err != nil {
		return Book{}, err
	}
	return Book{
		ID:          uuid.New().String(),
		Title:       title,
		Author:      author,
		Year:        year,
		IsAvailable: true,
		AddedAt:     time.Now(),
		BorrowedAt:  nil,
	}, nil
}

func (b *Book) Borrow() {
	borrowedTime := time.Now()

	b.BorrowedAt = &borrowedTime
	b.IsAvailable = false
}

func (b *Book) Return() {
	b.BorrowedAt = nil
	b.IsAvailable = true
}
