package database

import (
	"MyBook/repository"

	"gorm.io/gorm"
)

func MigrateBook(db *gorm.DB) {
	db.AutoMigrate(&repository.Book{})
}
