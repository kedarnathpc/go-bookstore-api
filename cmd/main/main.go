package main

import (
	"bookstore/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//create a router
	r := mux.NewRouter()
	//pass the router to handle the functions
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	// start the server
	fmt.Println("Starting server at port :9010 ...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
