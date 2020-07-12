package main

import (
	"fmt"
	"net/http"
)

func Allusers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "All users end point")
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "new user end point")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Delete user end point")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "update user end point")
}
