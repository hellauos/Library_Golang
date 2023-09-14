package category

import (
	"pustaka-api/book"
)

type Category struct {
	ID    uint
	Name  string
	Books []book.Book
}
