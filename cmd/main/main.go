package main

import (
	"go-lang-crud-api-mysql/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBooksStoreRoutes(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
