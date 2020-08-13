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
		article{Title: "docker compose", Desc: "compose des", Content: "content"},
		article{Title: "golang", Desc: "go with rest", Content: "vicky content"},
	}
	fmt.Println("All articles end point")
	json.NewEncoder(w).Encode(articles)

}

// func testArticles(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(w, "post method")
// }

// func homepage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(w, "Docker compose page")

// }

// func handleRequests() {

// 	http.HandleFunc("/", homepage)
// 	http.HandleFunc("/articles", allArticles)
// 	http.HandleFunc("/articless", testArticles)

// }

func main() {
	handleRequests()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
