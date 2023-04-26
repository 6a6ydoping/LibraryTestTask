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

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author

	db.DB.Preload("Books").Find(&authors)

	authorsJSON, err := json.Marshal(authors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(authorsJSON)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&author)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}

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

func GetAuthorBooks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var author models.Author
	if result := db.DB.Preload("Books").First(&author, authorID); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author.Books)
}
