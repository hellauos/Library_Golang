package book

import (
	"pustaka-api/loan"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	Description string
	Location    string
	Price       int
	Stock       int
	CategoryID  uint
	Loans       []loan.Loan
}
