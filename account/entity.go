package account

import (
	"pustaka-api/loan"
)

type Account struct {
	ID       uint
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Loans    []loan.Loan
	RolesId  uint
}
