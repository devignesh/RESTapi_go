package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]
	fmt.Println("Hello, playground")
	fmt.Printf("myname is: %s\n", name)
}
