package book

import (
	"time"
)

type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	Location    string
	Stock       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// Loan        []loan.Loan
	Category int
}
