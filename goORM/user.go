package main

import (
	"fmt"
	"net/http"
)

func Allusers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "All users end point")
}
