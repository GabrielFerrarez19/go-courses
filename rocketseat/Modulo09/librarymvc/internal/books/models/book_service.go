package model

type BookService interface {
	CreateBook(user *Book) error
	GetBookByID(id int64) (*Book, error)
	ListBook() ([]*Book, error)
	UpdateBook(user *Book, id int64) (*Book, error)
	DeleteBook(id int64) error
}
