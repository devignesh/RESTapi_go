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
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handlerequests()
}
