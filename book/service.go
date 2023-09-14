package book

type Service interface {
	FindAll() ([]Book, error)
	FindAllBooksByUser(UserID uint) ([]Book, error)
	FindById(ID int) (Book, error)
	Create(bookRequest BookRequest, UserId uint) (Book, error)
	Update(ID int, bookRequest BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindById(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest, UserId uint) (Book, error) {
	book := Book{
		Title:       bookRequest.Title,
		Price:       bookRequest.Price,
		Description: bookRequest.Description,
		// Rating:      bookRequest.Rating,
		// UserId:      UserId,
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	book, err := s.repository.FindById(ID)

	if bookRequest.Title != "" {
		book.Title = bookRequest.Title
	}
	if bookRequest.Price != 0 {
		book.Price = bookRequest.Price
	}
	if bookRequest.Description != "" {
		book.Description = bookRequest.Description
	}
	// if bookRequest.Rating != 0 {
	// 	book.Rating = bookRequest.Rating
	// }

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	_, err = s.repository.Delete(book)
	return book, err
}

func (s *service) FindAllBooksByUser(UserID uint) ([]Book, error) {
	books, err := s.repository.FindAllBooksByUser(UserID)
	return books, err
}
