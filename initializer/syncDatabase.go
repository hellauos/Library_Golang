package initializer

import (
	"pustaka-api/book"
	"pustaka-api/category"
	"pustaka-api/loan"
	"pustaka-api/roles"
	"pustaka-api/user"

	"gorm.io/gorm"
)

func SyncDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(roles.Roles{}, user.User{}, category.Category{}, book.Book{}, loan.Loan{})
	return err
}
