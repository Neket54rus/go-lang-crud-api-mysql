package routes

import (
	"go-lang-crud-api-mysql/pkg/controllers"

	"github.com/gorilla/mux"
)

func RegisterBooksStoreRoutes(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBookById).Methods("DELETE")
}
