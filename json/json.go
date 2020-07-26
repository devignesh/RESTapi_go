package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

func main() {
	fmt.Println("Hello Json")
	bk := Book{Title: "my book", Author: "vicky",Desc:"test book"}
	fmt.Printf("%+v\n", bk)

	jsonmar, err := json.Marshal(bk)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonmar))
}
