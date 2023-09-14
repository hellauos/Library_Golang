package book

import (
	"pustaka-api/loan"
)

type Book struct {
	ID          uint
	Title       string
	Description string
	Location    string
	Price       int
	Stock       int
	CategoryID  uint
	Loans       []loan.Loan
}
