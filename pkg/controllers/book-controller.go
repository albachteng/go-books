package controllers

import (
	"encoding/json"
	"fmt"
	"go-books/pkg/models"
	"go-books/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	// retrieve the book documents
	newBooks := models.GetAllBooks()
	// convert the book documents to json
	response, _ := json.Marshal(newBooks)
	// send it back in the response
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	// get the url params
	vars := mux.Vars(r)
	// we're interested in the ID specifically
	bookId := vars["bookId"]
	// parse the ID string as an integer
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// retrieve the book document by ID
	bookDetails, _ := models.GetBookById(ID)
	// parse the document into JSON
	response, _ := json.Marshal(bookDetails)
	// send it all back
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// create a book type that is empty for now
	CreateBook := &models.Book{}
	// parse the request body onto the empty book type we just declared
	utils.ParseBody(r, CreateBook)
	// ?
	b := CreateBook.CreateBook()
	response, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	response, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// empty book to update
	var updateBook = &models.Book{}
	// details of the book that we will use to update
	utils.ParseBody(r, updateBook)
	// variables from the url
	vars := mux.Vars(r)
	// the ID of the book to update
	bookId := vars["bookId"]
	// parse the string into a number
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// retrieve the actual book document we are updating
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	// save it and send the book in our response as JSON
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
