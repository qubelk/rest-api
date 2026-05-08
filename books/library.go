package books

type Library struct {
	// key == ID
	books map[string]Book
}

func bookExists(l *Library, id string) bool {
	if _, exists := l.books[id]; exists {
		return true
	}

	return false
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

func (l *Library) AddBook(book Book) error {
	if bookExists(l, book.ID) {
		return BookAlreadyExists
	}
	l.books[book.ID] = book
	return nil
}

func (l *Library) DeleteBook(id string) error {
	if !bookExists(l, id) {
		return BookNotExists
	}
	delete(l.books, id)
	return nil
}

func (l *Library) GetAllBooks() map[string]Book {
	tmp := make(map[string]Book, len(l.books))
	for k, v := range tmp {
		tmp[k] = v
	}

	return tmp
}

func (l *Library) GetAllBorrowedBooks() map[string]Book {
	tmp := make(map[string]Book)
	for k, v := range tmp {
		if !v.IsAvailable {
			tmp[k] = v
		}
	}

	return tmp
}

func (l *Library) GetBookByID(id string) (Book, error) {
	if !bookExists(l, id) {
		return Book{}, BookNotExists
	}
	return l.books[id], nil
}

func (l *Library) GetBookByAuthor(author string) (Book, error) {
	for _, v := range l.books {
		if v.Author == author {
			return v, nil
		}
	}
	return Book{}, BookNotExists
}

func (l *Library) BorrowBook(id string) error {
	if !bookExists(l, id) {
		return BookNotExists
	}

	b := l.books[id]
	b.Borrow()
	b.IsAvailable = false
	l.books[id] = b
	return nil
}

func (l *Library) ReturnBook(id string) error {
	if !bookExists(l, id) {
		return BookNotExists
	}

	b := l.books[id]
	b.Return()
	b.IsAvailable = true
	l.books[id] = b
	return nil
}
