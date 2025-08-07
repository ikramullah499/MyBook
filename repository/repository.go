package repository

import (
	"gorm.io/gorm"
)

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type BookRepository interface {
	Create(book *Book) error
	FindAll() ([]Book, error)
	FindByID(id string) (*Book, error)
	Delete(id string) error
}

type bookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepositoryImpl{db: db}
}

func (r *bookRepositoryImpl) Create(book *Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepositoryImpl) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepositoryImpl) FindByID(id string) (*Book, error) {
	var book Book
	err := r.db.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&Book{}, id).Error
}
