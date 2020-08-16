package main

import (
	"fmt"
	"net/http"
	"os"
)

func roothandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Awesome site in Go!</h1><br>")
	fmt.Fprintf(w, "<a href='/host/'>Host info</a><br>")
}

func hosthandler(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "<h1>HOSTNAME: %s</h1><br>", name)
	fmt.Fprintf(w, "<h1>ENVIRONMENT VARS:</h1><br>")
	fmt.Fprintf(w, "<ul>")

	for _, evar := range os.Environ() {
		fmt.Fprintf(w, "<li>%s</li>", evar)
	}
	fmt.Fprintf(w, "</ul>")
}

func main() {

	http.HandleFunc("/", roothandler)
	http.HandleFunc("/host/", hosthandler)
	http.ListenAndServe(":8081", nil)
}
