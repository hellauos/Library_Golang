package category

import (
	"pustaka-api/book"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string
	Books []book.Book
}
