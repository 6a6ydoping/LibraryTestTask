package controllers

import (
	"encoding/json"
	"github.com/6a6ydoping/LibraryTestTask/db"
	"github.com/6a6ydoping/LibraryTestTask/models"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	result := db.DB.Find(&books)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to retrieve books"))
		return
	}
	booksJSON, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to marshal books to JSON"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(booksJSON)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to parse request body"))
		return
	}
	author := models.Author{ID: book.AuthorID}
	result := db.DB.First(&author)
	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Author not found"))
		return
	}
	newBook := models.Book{
		AuthorID: book.AuthorID,
		Title:    book.Title,
		Genre:    book.Genre,
		ISBN:     book.ISBN,
	}
	result = db.DB.Create(&newBook)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to create book"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}
