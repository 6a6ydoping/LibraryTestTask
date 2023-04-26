package controllers

import (
	"encoding/json"
	"github.com/6a6ydoping/LibraryTestTask/db"
	"github.com/6a6ydoping/LibraryTestTask/models"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func CreateReader(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var reader models.Reader
	err = json.Unmarshal(body, &reader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.DB.Create(&reader)
}

func GetAllMembers(w http.ResponseWriter, r *http.Request) {
	var readers []models.Reader
	db.DB.Find(&readers)
	response, err := json.Marshal(readers)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func BorrowBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	readerID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var requestData struct {
		BookID uint `json:"book_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var reader models.Reader
	if err := db.DB.Preload("BorrowedBooks").First(&reader, uint(readerID)).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var book models.Book
	if err := db.DB.First(&book, requestData.BookID).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for _, borrowedBook := range reader.BorrowedBooks {
		if borrowedBook.ID == book.ID {
			w.WriteHeader(http.StatusConflict)
			return
		}
	}

	if err := db.DB.Model(&reader).Association("BorrowedBooks").Append(&book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reader); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetReaderBorrowedBooks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	readerID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var reader models.Reader
	if err := db.DB.Preload("BorrowedBooks").First(&reader, readerID).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reader.BorrowedBooks)
}
