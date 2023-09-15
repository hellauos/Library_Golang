package book

type GetBookByTitleCategoryRequest struct {
	Title    string `json:"title"`
	Category string `json:"category"`
}

type BookRequest struct {
	Title       string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	Price       int    `json:"price" binding:"required,number"`
	Stock       int    `binding:"required,number"`
	CategoryID  uint   `binding:"required,number"`
}
