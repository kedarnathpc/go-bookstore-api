package routes

import (
	"github.com/gorilla/mux"

	//include the controllers
	"bookstore/pkg/controllers"
)

// create a function with all the routes
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/create", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
