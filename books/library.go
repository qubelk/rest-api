package books

import (
	"maps"
	"rest-api/sins"
	"sync"

	"github.com/google/uuid"
)

type Library struct {
	// key == ID
	books map[string]Book
	mtx   sync.RWMutex
}

func NewLibrary() Library {
	return Library{
		books: make(map[string]Book),
	}
}

func NewLibraryWithCapacity(capacity uint64) Library {
	return Library{
		books: make(map[string]Book, capacity),
	}
}

func (l *Library) AddBook(book *Book) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	book.ID = uuid.New().String()
	if _, exists := l.books[book.ID]; exists {
		return sins.BookAlreadyExists
	}

	l.books[book.ID] = *book
	return nil
}

func (l *Library) DeleteBook(id string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, exists := l.books[id]; !exists {
		return sins.BookNotExists
	}

	delete(l.books, id)
	return nil
}

func (l *Library) GetAllBooks() map[string]Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tmp := make(map[string]Book, len(l.books))
	maps.Copy(tmp, l.books)

	return tmp
}

func (l *Library) GetAllAvailableBooks() map[string]Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tmp := make(map[string]Book)
	for k, v := range l.books {
		if v.IsAvailable {
			tmp[k] = v
		}
	}

	return tmp
}

func (l *Library) GetAllBorrowedBooks() map[string]Book {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	tmp := make(map[string]Book)
	for k, v := range l.books {
		if !v.IsAvailable {
			tmp[k] = v
		}
	}

	return tmp
}

func (l *Library) GetBookByID(id string) (Book, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	if _, exists := l.books[id]; !exists {
		return Book{}, sins.BookNotExists
	}

	return l.books[id], nil
}

func (l *Library) BorrowBook(id string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	b, exists := l.books[id]
	if !exists {
		return sins.BookNotExists
	}

	b.Borrow()
	l.books[id] = b
	return nil
}

func (l *Library) ReturnBook(id string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	b, exists := l.books[id]
	if !exists {
		return sins.BookNotExists
	}

	b.Return()
	l.books[id] = b
	return nil
}
