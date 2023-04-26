package routes

import (
	"github.com/6a6ydoping/LibraryTestTask/controllers"
	"github.com/gorilla/mux"
)

var Router *mux.Router

func RegisterRoutes() {
	Router = mux.NewRouter()
	Router.HandleFunc("/authors", controllers.GetAllAuthors).Methods("GET")
	Router.HandleFunc("/authors", controllers.CreateAuthor).Methods("POST")

	Router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	Router.HandleFunc("/books", controllers.CreateBook).Methods("POST")

	Router.HandleFunc("/members", controllers.GetAllMembers).Methods("GET")
	Router.HandleFunc("/members", controllers.CreateReader).Methods("POST")

	Router.HandleFunc("/members/{id}/books", controllers.BorrowBook).Methods("POST")
	Router.HandleFunc("/authors/{id}/books", controllers.GetAuthorBooks).Methods("GET")

	Router.HandleFunc("/authors/{id}/books", controllers.GetAuthorBooks).Methods("GET")
	Router.HandleFunc("/members/{id}/books", controllers.GetReaderBorrowedBooks).Methods("GET")
}
