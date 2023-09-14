package account

import (
	"pustaka-api/loan"
	"pustaka-api/roles"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	name     string
	Email    string `gorm:"unique"`
	Password string
	Loans    []loan.Loan
	Roles    roles.Roles
}
