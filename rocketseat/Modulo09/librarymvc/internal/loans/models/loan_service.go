package models

type LoansService interface {
	CreateLoan(bookID, userID int64) (*Loan, error)
	ReturnBook(loanID int64) error
	GetLoanByID(id int64) (*Loan, error)
	ListLoan() ([]*Loan, error)
	ListLoanByUserID(userID int64) ([]*Loan, error)
}
