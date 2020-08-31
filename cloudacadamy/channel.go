package main

import "fmt"

func main() {
	mychan := make(chan string)

	go func() {
		mychan <- "vicky"
		mychan <- "kumar"
		mychan <- "gopher"
	}()

	msg1 := <-mychan
	msg2 := <-mychan
	msg3 := <-mychan

	fmt.Println(msg1, msg2, msg3)
}
