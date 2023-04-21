package main

import (
	"github.com/6a6ydoping/LibraryTestTask/db"
	"github.com/6a6ydoping/LibraryTestTask/routes"
	"net/http"
)

func init() {
	db.Connect()
	db.SyncDB()
	routes.RegisterRoutes()
}

func main() {
	http.ListenAndServe("localhost:8080", routes.Router)
}
