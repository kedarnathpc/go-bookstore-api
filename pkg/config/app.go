// to return a variable called db which will help other files to interact with the db
package config

import (
	//import the required packages
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// create a DB variable
var (
	db *gorm.DB
)

// create a function to open your connection to the database
func Connect() {
	d, err := gorm.Open("mysql", "user:root@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
