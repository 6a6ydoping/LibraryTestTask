package controllers

import (
	"encoding/json"
	"github.com/6a6ydoping/LibraryTestTask/db"
	"github.com/6a6ydoping/LibraryTestTask/models"
	"github.com/gorilla/mux"
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
