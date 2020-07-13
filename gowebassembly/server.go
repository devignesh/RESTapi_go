package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", "8080", "listen address")
	dir    = flag.String("dir", ".", "Directory to server")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q..", *listen)
	log.Fatal(http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir))))
}
