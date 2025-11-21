package services

import (
	"time"

	"librarymvc/internal/books/models"
)

type BookService struct {
	repo models.BookRepository
}

func NewBookService(repo models.BookRepository) models.BookService {
	return &BookService{
		repo: repo,
	}
}

// CreateBook implements models.BookService.
func (b *BookService) CreateBook(book *models.Book) error {
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	return b.repo.CreateBook(book)
}

// DeleteBook implements models.BookService.
func (b *BookService) DeleteBook(id int64) error {
	return b.repo.DeleteBook(id)
}

// GetBookByID implements models.BookService.
func (b *BookService) GetBookByID(id int64) (*models.Book, error) {
	return b.repo.GetBookByID(id)
}

// ListBook implements models.BookService.
func (b *BookService) ListBook() ([]*models.Book, error) {
	return b.repo.ListBook()
}

// UpdateBook implements models.BookService.
func (b *BookService) UpdateBook(book *models.Book, id int64) (*models.Book, error) {
	book.UpdatedAt = time.Now()

	return b.repo.UpdateBook(book, id)
}
