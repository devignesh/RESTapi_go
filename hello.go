package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const defaultAddr = ":8080"


func main() {
	addr := defaultAddr
	
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}

	log.Printf("server starting to listen on %s", addr)
	http.HandleFunc("/", home)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server listen error: %+v", err)
	}
}


func home(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request: %s %s", r.Method, r.URL.Path)
	fmt.Fprintf(w, "Hello, world!")
}
