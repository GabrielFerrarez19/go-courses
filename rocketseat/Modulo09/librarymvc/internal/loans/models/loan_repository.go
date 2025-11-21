package models

type LoansRepository interface {
	CreateLoan(loan *Loan) error
	UpdatedLoan(loan *Loan) error
	ReturnBook(loanID int64) error
	GetLoanByID(id int64) (*Loan, error)
	ListLoan() ([]*Loan, error)
	ListActiveLoanByUserID(userID int64) ([]*Loan, error)
}
