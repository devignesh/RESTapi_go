package main

import (
	"fmt"
	"log"
	"net/http"
)

func kloud(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Kloudone")

}

func setuproutes() {
	http.HandleFunc("/", kloud)
}

func main() {
	fmt.Println("Kloudone with docker test")
	setuproutes()
	log.Fatal(http.ListenAndServe(":8080", nil))

}
