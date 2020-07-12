package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func sample(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "sample function")
}

func handlerequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", sample).Methods("GET")
	router.HandleFunc("/users", Allusers).Methods("GET")
	router.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	router.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handlerequests()
}
