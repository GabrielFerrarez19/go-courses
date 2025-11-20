package models

import "time"

type Loan struct {
	ID         int64     `json:"ID"`
	UserID     int64     `json:"userID"`
	BookID     int64     `json:"bookID"`
	BorrowedAt time.Time `json:"borrowedAt"`
	ReturnedAt time.Time `json:"returnedAt"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
