package middlewares

import (
	"encoding/json"
	"github.com/6a6ydoping/LibraryTestTask/db"
	"github.com/6a6ydoping/LibraryTestTask/models"
	"io"
	"net/http"
)

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var author models.Author
	err = json.Unmarshal(body, &author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.DB.Create(&author)
}
