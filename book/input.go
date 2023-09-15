package book

type GetBookByTitleCategoryRequest struct {
	Title    string `json:"title"`
	Category string `json:"category"`
}

// type BookRequest struct {
// 	Title       string `binding:"required"`
// 	Price       int    `json:"price" binding:"required,number"`
// 	Description string `binding:"required"`
// 	Rating      int    `binding:"required,number"`
// }

// type UpdateBookRequest struct {
// 	Title       string
// 	Price       int
// 	Description string
// 	Rating      int
// }
