package repository

import (
	"errors"
	"sync"

	"librarymvc/internal/books/models"
)

type BookRepository struct {
	book   map[int64]*models.Book
	mu     sync.RWMutex
	nextId int64
}

func NewBookRepository() models.BookRepository {
	return &BookRepository{
		book:   map[int64]*models.Book{},
		nextId: 1,
	}
}

// CreateBook implements models.BookRepository.
func (b *BookRepository) CreateBook(book *models.Book) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	book.ID = b.nextId
	b.nextId++

	b.book[book.ID] = book
	return nil
}

// DeleteBook implements models.BookRepository.
func (b *BookRepository) DeleteBook(id int64) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	_, exists := b.book[id]
	if !exists {
		return errors.New("user not found")
	}

	delete(b.book, id)

	return nil
}

// GetBookByID implements models.BookRepository.
func (b *BookRepository) GetBookByID(id int64) (*models.Book, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	book, exists := b.book[id]
	if !exists {
		return nil, errors.New("book not found")
	}
	return book, nil
}

// ListBook implements models.BookRepository.
func (b *BookRepository) ListBook() ([]*models.Book, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	books := make([]*models.Book, 0, len(b.book))

	for _, book := range b.book {
		books = append(books, book)
	}
	return books, nil
}

// UpdateBook implements models.BookRepository.
func (b *BookRepository) UpdateBook(book *models.Book, id int64) (*models.Book, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	_, exists := b.book[id]
	if !exists {
		return nil, errors.New("book not found")
	}

	b.book[book.ID] = book

	return book, nil
}
