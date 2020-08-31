package main

import "fmt"
import "time"

func main() {
	channel := make(chan string)

	go func(channel chan string) {
		time.Sleep(5 * time.Second)
		channel <- "vickykumar"

	}(channel)

	select {
	case msg1 := <-channel:
		fmt.Println(msg1)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout!")
	}
}
