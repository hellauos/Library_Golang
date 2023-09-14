package initializer

import (
	"pustaka-api/account"
	"pustaka-api/book"
	"pustaka-api/category"
	"pustaka-api/loan"
	"pustaka-api/role"

	"gorm.io/gorm"
)

func SyncDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(book.Book{}, account.Account{}, loan.Loan{}, category.Category{}, role.Role{})
	return err
}
