package repository

import (
	"errors"
	"sync"
	"time"

	"librarymvc/internal/loans/models"
)

type LoansRepository struct {
	loan   map[int64]*models.Loan
	mu     sync.RWMutex
	nextId int64
}

func NewLoansRepository() models.LoansRepository {
	return &LoansRepository{
		loan:   make(map[int64]*models.Loan),
		nextId: 1,
	}
}

// CreateLoan implements models.LoansRepository.
func (l *LoansRepository) CreateLoan(loan *models.Loan) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	loan.ID = l.nextId
	l.nextId++

	l.loan[loan.ID] = loan
	return nil
}

// GetLoanByID implements models.LoansRepository.
func (l *LoansRepository) GetLoanByID(id int64) (*models.Loan, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	loan, exists := l.loan[id]
	if !exists {
		return nil, errors.New("loan not found")
	}

	return loan, nil
}

// ListActiveLoanByUserID implements models.LoansRepository.
func (l *LoansRepository) ListActiveLoanByUserID(userID int64) ([]*models.Loan, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	loans := make([]*models.Loan, 0, len(l.loan))

	for _, loan := range l.loan {
		if loan.UserID == userID && loan.Status == "active" {
			loans = append(loans, loan)
		}
	}

	return loans, nil
}

// ListLoan implements models.LoansRepository.
func (l *LoansRepository) ListLoan() ([]*models.Loan, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	loans := make([]*models.Loan, 0, len(l.loan))

	for _, loan := range l.loan {
		loans = append(loans, loan)
	}

	return loans, nil
}

// ReturnBook implements models.LoansRepository.
func (l *LoansRepository) ReturnBook(loanID int64) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	loan, exists := l.loan[loanID]
	if !exists {
		return errors.New("loan not found")
	}

	loan.Status = "returned"
	loan.ReturnedAt = time.Now()

	return nil
}

// UpdatedLoan implements models.LoansRepository.
func (l *LoansRepository) UpdatedLoan(loan *models.Loan) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	_, exists := l.loan[loan.ID]
	if !exists {
		return errors.New("loan not found")
	}

	l.loan[loan.ID] = loan

	return nil
}
