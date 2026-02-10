package controllers

import (
	"encoding/json"
	"fmt"
	"go-lang-crud-api-mysql/pkg/models"
	"go-lang-crud-api-mysql/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(res http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()

	r, _ := json.Marshal(newBooks)
	res.Header().Set("Content-type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r)
}

func GetBookById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	bookId, err := strconv.ParseInt(id, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails, _ := models.GetBookById(bookId)

	r, _ := json.Marshal(bookDetails)

	res.Header().Set("Content-type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r)
}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(req, CreateBook)

	b := CreateBook.CreateBook()

	r, _ := json.Marshal(b)

	res.WriteHeader(http.StatusOK)
	res.Write(r)
}

func UpdateBookById(res http.ResponseWriter, req *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(req, updateBook)

	vars := mux.Vars(req)
	id := vars["id"]

	bookId, err := strconv.ParseInt(id, 0, 0)

	if err != nil {
		fmt.Println("Error while pasring")
	}

	bookDetails, db := models.GetBookById(bookId)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	r, _ := json.Marshal(bookDetails)

	res.Header().Set("Content-type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r)
}

func DeleteBookById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	bookId, err := strconv.ParseInt(id, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	book := models.DeleteBookById(bookId)

	r, _ := json.Marshal(book)

	res.Header().Set("Content-type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r)
}
