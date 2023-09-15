package user

import (
	"pustaka-api/loan"
)

type User struct {
	ID       uint
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Loans    []loan.Loan
	RolesId  uint
}
