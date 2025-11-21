package models

type BookRepository interface {
	CreateBook(book *Book) error
	GetBookByID(id int64) (*Book, error)
	ListBook() ([]*Book, error)
	UpdateBook(book *Book, id int64) (*Book, error)
	DeleteBook(id int64) error
}
