package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func sample(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "sample function")
}

func handlerequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", sample).Methods("GET")

}
