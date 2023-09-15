package loan

import (
	"errors"
	"fmt"
	"pustaka-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	LoanBook(loanBookRequest LoanBookRequest) (string, error)
	ReturnBook(returnBookRequest ReturnBookRequest) (string, error)
	FindLoanData() ([]Loan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) LoanBook(loanBookRequest LoanBookRequest) (string, error) {
	// validate user
	var account Account
	err := r.db.First(&account, loanBookRequest.AccountId).Error
	if err != nil {
		return "", err
	}
	// check stock
	var book Book

	if err := r.db.First(&book, loanBookRequest.BookId).Error; err != nil {
		return "", err
	} else if book.Stock <= 0 {
		return "", errors.New("the book is empty")
	}

	// make object
	loan := Loan{
		DueDate:   helper.StringToDate(loanBookRequest.DueDate),
		Status:    1,
		BookId:    uint(loanBookRequest.BookId),
		AccountId: uint(loanBookRequest.AccountId),
	}
	// maximum loan
	var count int64
	if err := r.db.Model(&Loan{}).Where("account_id = ?", 1).Count(&count).Error; err != nil {
		return "", err
	} else if count > int64(5) {
		fmt.Println("exceed")
		return "", errors.New("maximum loan exceed")
	}
	// do loan
	err = r.db.Create(&loan).Error

	//decrease stock
	book.Stock = book.Stock - 1

	// Save the updated field back to the database
	if err := r.db.Save(&book).Error; err != nil {
		return "", err
	}

	return "Loan success", nil
}

func (r *repository) ReturnBook(returnBookRequest ReturnBookRequest) (string, error) {
	// loan status turn to 0
	var loan Loan

	if err := r.db.First(&loan, returnBookRequest.LoanId).Error; err != nil {
		return "", err
	}
	loan.Status = 0
	if err := r.db.Save(&loan).Error; err != nil {
		return "", err
	}

	// update stock book
	var book Book
	if err := r.db.First(&book, loan.BookId).Error; err != nil {
		return "", err
	}
	book.Stock = book.Stock + 1
	if err := r.db.Save(&book).Error; err != nil {
		return "", err
	}

	return "Return success", nil
}

func (r *repository) FindLoanData() ([]Loan, error) {
	var loan []Loan
	err := r.db.Find(&loan).Error
	return loan, err

}
