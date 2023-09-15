package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	FindByTitle(title string) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, bookRequest BookRequest) (Book, error)
	UpdateByTitle(title string, bookRequest BookRequest) (Book, error)
	Delete(ID int) (Book, error)
	DeleteByTitle(title string) (Book, error)
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

func (s *service) FindByTitle(title string) (Book, error) {
	book, err := s.repository.FindByTitle(title)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	book := Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Location:    bookRequest.Location,
		Price:       bookRequest.Price,
		Stock:       bookRequest.Stock,
		CategoryID:  bookRequest.CategoryID,
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	book, err := s.repository.FindById(ID)

	if bookRequest.Title != "" {
		book.Title = bookRequest.Title
	}
	if bookRequest.Description != "" {
		book.Description = bookRequest.Description
	}
	if bookRequest.Location != "" {
		book.Location = bookRequest.Location
	}
	if bookRequest.Price != 0 {
		book.Price = bookRequest.Price
	}
	if bookRequest.Stock != 0 {
		book.Stock = bookRequest.Stock
	}
	if bookRequest.CategoryID != 0 {
		book.CategoryID = bookRequest.CategoryID
	}
	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) UpdateByTitle(title string, bookRequest BookRequest) (Book, error) {
	book, err := s.repository.FindByTitle(title)

	if bookRequest.Title != "" {
		book.Title = bookRequest.Title
	}
	if bookRequest.Description != "" {
		book.Description = bookRequest.Description
	}
	if bookRequest.Location != "" {
		book.Location = bookRequest.Location
	}
	if bookRequest.Price != 0 {
		book.Price = bookRequest.Price
	}
	if bookRequest.Stock != 0 {
		book.Stock = bookRequest.Stock
	}
	if bookRequest.CategoryID != 0 {
		book.CategoryID = bookRequest.CategoryID
	}
	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	_, err = s.repository.Delete(book)
	return book, err
}

func (s *service) DeleteByTitle(title string) (Book, error) {
	book, err := s.repository.FindByTitle(title)
	_, err = s.repository.Delete(book)
	return book, err
}
