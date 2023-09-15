package initializer

import (
	"pustaka-api/book"
	"pustaka-api/category"
	"pustaka-api/loan"
	"pustaka-api/roles"
	"pustaka-api/account"

	"gorm.io/gorm"
)

func SyncDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(roles.Roles{}, account.Account{}, category.Category{}, book.Book{}, loan.Loan{})
	return err
}
