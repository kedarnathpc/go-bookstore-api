package controllers

import (
	"bookstore/pkg/models"
	"bookstore/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// create a new book of type Book in from the models dir
var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {

	//get all the books from the database through models
	newBooks := models.GetAllBooks()

	//marshal it into the json format
	res, _ := json.Marshal(newBooks)

	//set the content header
	w.Header().Set("Content-Type", "pkglication/json")

	//write the status of the  operation
	w.WriteHeader(http.StatusOK)

	// return the response
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {

	//get the book id from the request
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	//convert the string id into integer id
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	//get the book details from the database through models
	bookDetails, _ := models.GetBookByID(ID)

	//create a response
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	//create a book var of type Book{} from the models
	CreateBook := &models.Book{}

	//convert the request into database understanding data
	utils.ParseBody(r, CreateBook)

	b := CreateBook.CreateBook()

	//create the response
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	//delete the book by calling delete model function from the models
	book := models.DeleteBook(ID)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	//create a book of type Book in the models
	var updateBook = &models.Book{}

	//parse the request body into the created variable
	utils.ParseBody(r, updateBook)

	//extract the book if
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	//get the book from the database
	bookDetails, db := models.GetBookByID(ID)

	// change the values of the book
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	// save the changes values into the database
	db.Save(&bookDetails)

	//create a response
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
