package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://www.google.com",
	"https://www.gopherhut.com",
	"https://www.github.com",
	"https://www.linkedin.com"
}

func status(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			response, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w, "%+v\n", err)
			}
			fmt.Fprintf(w, "%+v\n", response)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func main() {
	fmt.Println("wait group example")

	http.HandleFunc("/", status)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
