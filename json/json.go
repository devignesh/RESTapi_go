package main

import (
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	fmt.Println("Hello Json")
	bk := Book{Title: "my book", Author: "vicky"}
	fmt.Printf("%+v\n", bk)
}
