package loan

import "fmt"

type Service interface {
	LoanBook(loanRequest LoanBookRequest) (string, error)
ReturnBook(returnBookRequest ReturnBookRequest) (string, error)
	FindLoanData() ([]Loan, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) LoanBook(loanBook LoanBookRequest) (string, error) {
	book, err := s.repository.LoanBook(loanBook)
	fmt.Println("service", err)
	return book, err
}

func (s *service) FindLoanData() ([]Loan, error) {
	loan, err := s.repository.FindLoanData()
	return loan, err
}

func (s *service) ReturnBook(returnBookRequest ReturnBookRequest) (string, error) {
	book, err := s.repository.ReturnBook(returnBookRequest)
	return book, err
}
