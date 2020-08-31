package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
	}("vicky")

	wg.Add(1)
	go func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
	}("kumar")

	wg.Wait()
}
