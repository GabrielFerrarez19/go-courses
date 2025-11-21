package services

import (
	"errors"
	"time"

	bookModel "librarymvc/internal/books/models"
	loanModel "librarymvc/internal/loans/models"
	userModel "librarymvc/internal/users/models"
)

type LoanService struct {
	repo        loanModel.LoansRepository
	bookService bookModel.BookService
	userService userModel.UserService
}

func NewLoanService(repo loanModel.LoansRepository, BookService bookModel.BookService, UserService userModel.UserService) loanModel.LoansService {
	return &LoanService{
		repo:        repo,
		bookService: BookService,
		userService: UserService,
	}
}

// CreateLoan implements models.LoansService.
func (l *LoanService) CreateLoan(bookID int64, userID int64) (*loanModel.Loan, error) {
	book, err := l.bookService.GetBookByID(bookID)
	if err != nil {
		return nil, err
	}

	if book.Quantity <= 0 {
		return nil, errors.New("boo is not available")
	}

	_, err = l.userService.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	loans, err := l.ListActiveLoanByUserID(userID)
	if err != nil {
		return nil, err
	}
	if len(loans) > 0 {
		return nil, errors.New("user has active loans")
	}

	loan := &loanModel.Loan{
		UserID:     userID,
		BookID:     bookID,
		BorrowedAt: time.Now(),
		Status:     "active",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err = l.repo.CreateLoan(loan)
	if err != nil {
		return nil, err
	}

	book.Quantity--

	_, err = l.bookService.UpdateBook(book, bookID)
	if err != nil {
		return nil, err
	}

	return loan, nil
}

// GetLoanByID implements models.LoansService.
func (l *LoanService) GetLoanByID(id int64) (*loanModel.Loan, error) {
	return l.repo.GetLoanByID(id)
}

// ListLoan implements models.LoansService.
func (l *LoanService) ListLoan() ([]*loanModel.Loan, error) {
	return l.repo.ListLoan()
}

// ListLoanByUserID implements models.LoansService.
func (l *LoanService) ListActiveLoanByUserID(userID int64) ([]*loanModel.Loan, error) {
	return l.repo.ListActiveLoanByUserID(userID)
}

// ReturnBook implements models.LoansService.
func (l *LoanService) ReturnBook(loanID int64) error {
	loan, err := l.GetLoanByID(loanID)
	if err != nil {
		return err
	}

	if loan.Status == "returned" {
		return errors.New("book already returned")
	}

	loan.Status = "returned"
	loan.UpdatedAt = time.Now()
	loan.ReturnedAt = time.Now()

	if err := l.repo.UpdatedLoan(loan); err != nil {
		return err
	}

	bookID := loan.BookID

	book, err := l.bookService.GetBookByID(bookID)
	if err != nil {
		return err
	}

	book.Quantity++

	_, err = l.bookService.UpdateBook(book, bookID)
	if err != nil {
		return err
	}

	return nil
}
