package book

func ConvertToBookResponse(b Book) BookResponse {
	return BookResponse{
		ID:          int(b.ID),
		Title:       b.Title,
		Description: b.Description,
		Location:    b.Location,
		Price:       b.Price,
		Stock:       b.Stock,
		CategoryID:  b.CategoryID,

		// Price:       b.Price,
	}
}
