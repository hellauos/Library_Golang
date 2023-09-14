package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	FindAllBooksByUser(UserID uint) ([]Book, error)
	Create(book Book) (Book, error)
	Update(book Book) (Book, error)
	Delete(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Book, error) {
	var books []Book

	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) FindById(ID int) (Book, error) {
	var book Book

	err := r.db.First(&book, ID).Error

	return book, err
}

func (r *repository) Create(book Book) (Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) Update(book Book) (Book, error) {
	err := r.db.Save(&book).Error
	return book, err
}

func (r *repository) Delete(book Book) (Book, error) {
	err := r.db.Delete(&book).Error
	return book, err
}

func (r *repository) FindAllBooksByUser(UserID uint) ([]Book, error) {
	var books []Book

	err := r.db.Where("user_id = ?", UserID).Find(&books).Error

	return books, err
}
