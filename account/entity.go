package account

import (
	"pustaka-api/loan"
	"pustaka-api/role"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	name     string
	Email    string `gorm:"unique"`
	Password string
	Loans    []loan.Loan
	Role     role.Role
}
