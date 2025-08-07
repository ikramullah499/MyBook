package database

import (
	"github.com/ikramullah499/book_api/repository"
	"gorm.io/gorm"
)

func MigrateBook(db *gorm.DB) {
	db.AutoMigrate(&repository.Book{})
}
