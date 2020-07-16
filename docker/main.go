package main

import (
	"fmt"
	"log"
	"net/http"
)

func docker(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello Docker")
}

func main() {
	fmt.Println("golang with Docker")
	http.HandleFunc("/", docker)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
