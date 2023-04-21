package db

import (
	"github.com/6a6ydoping/LibraryTestTask/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := "host=localhost user=postgres password=123 dbname=book-store port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database...")
	}
}

func SyncDB() {
	DB.AutoMigrate(&models.Book{}, &models.Reader{}, &models.Author{}, &models.BorrowedBooks{})
}
