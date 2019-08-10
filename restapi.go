package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articals []article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articals{
		article{Title: "test title", Desc: "test description", Content: "Hi vicky"},
	}
	fmt.Println("All articles end point")
	json.NewEncoder(w).Encode(articles)

}

func testArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintlf(w, "test part of post method")
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sample page RESTapi")

}

func handleRequests() {

	myrouter := mux.NewRouter().StrictSlash(true)

	myrouter.HandleFunc("/", homepage)
	myrouter.HandleFunc("/articles", allArticles).Methods("GET")
	myrouter.HandleFunc("/articles", testArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myrouter))

}

func main() {
	handleRequests()
}
