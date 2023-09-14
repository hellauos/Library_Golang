package initializer

import (
	"pustaka-api/book"
	"pustaka-api/user"

	"gorm.io/gorm"
)

func SyncDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(book.Book{}, user.User{})
	return err
}
