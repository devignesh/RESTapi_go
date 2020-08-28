package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type article struct {
	id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articals []article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articals{
		article{id: "1", Title: "test title", Desc: "test description", Content: "Hi vicky"},
	}
	fmt.Println("All articles end point")
	json.NewEncoder(w).Encode(articles)

}

func singlearticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articals {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article article
	json.Unmarshal(reqBody, &article)

	article = append(Articals, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articals {
		if article.Id == id {
			article = append(Articals[:index], Articals[index+1:]...)
		}
	}

}

func testArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "test part of post method")
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Sample page RESTapi")

}

func handleRequests() {

	rot := mux.NewRouter().StrictSlash(true)

	rot.HandleFunc("/", homepage)
	rot.HandleFunc("/articles", allArticles).Methods("GET")
	rot.HandleFunc("/articlee", createNewArticle).Methods("POST")
	rot.HandleFunc("/articlee/{id}", singlearticle).Methods("GET")
	rot.HandleFunc("/articlee/{id}", deleteArticle).Methods("GET")
	rot.HandleFunc("/articlex", testArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", rot))

}

func main() {
	handleRequests()
}
