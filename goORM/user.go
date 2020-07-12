package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db gorm.DB
var err error

type User struct {
	gorm.Model

	Name  string
	email string
}

func IntialMigration() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func Allusers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "All users end point")
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "new user end point")

	fmt.Println(w, "All users end point")
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{name: name, email: email})

	fmt.Fprintf(w, "new user successfully")

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Delete user end point")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "update user end point")
}
