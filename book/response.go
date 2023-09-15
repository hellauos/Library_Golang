package book

type BookResponse struct {
	ID          int
	Title       string
	Description string
	Location    string
	Price       int
	Stock       int
	CategoryID  uint
}
