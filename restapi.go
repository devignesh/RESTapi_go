package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Sample page RESTapi")

}

func handleRequests() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func main() {
	handleRequests()
}
