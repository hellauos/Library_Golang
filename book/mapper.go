package book

func ConvertToBookResponse(b Book) BookResponse {
	return BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Rating:      b.Rating,
	}
}
