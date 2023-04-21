package routes

import (
	"github.com/6a6ydoping/BookManager/controllers"
	"github.com/gorilla/mux"
)

var Router *mux.Router

func RegisterRoutes() {
	Router = mux.NewRouter()
	Router.HandleFunc("/authors", controllers.CreateAuthor)
	Router.HandleFunc("/books", nil)
	Router.HandleFunc("/members", nil)
	Router.HandleFunc("/authors/{id}/books", nil)
	Router.HandleFunc("/members/{id}/books", nil)
}
