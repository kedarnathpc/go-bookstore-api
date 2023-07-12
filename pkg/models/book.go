package models

import (
	"bookstore/pkg/config"

	"github.com/jinzhu/gorm"
)

// create a database variable
var db *gorm.DB

// struct are based on models, which help to store data in database
type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// function to initialize the database
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// model functions to connect with the database from the controller functions
func (b *Book) CreateBook() *Book {
	//function in gorm to handle new record quires
	db.NewRecord(b)
	//function in gorm to create a new record
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
