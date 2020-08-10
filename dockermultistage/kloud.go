package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func kloudfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Welcome to the kloudone")
}

func routes() {
	router := mux.NewRouter()
	router.HandleFunc("/", kloudfunc)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	fmt.Println("welcome")
	routes()
	//log.Fatal(http.ListenAndServe(":8080", nil))

}
