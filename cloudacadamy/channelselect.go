package main

import (
	"fmt"
	"time"
)

func in(channel chan<- string, msg string) {
	channel <- msg
}

func out(channel <-chan string) {
	for {
		fmt.Println(<-channel)
	}
}

func main() {
	channel := make(chan string, 1)

	go out(channel)

	for i := 0; i < 10; i++ {
		in(channel, fmt.Sprintf("vickykumar - %d", i))
	}

	time.Sleep(time.Second * 10)
}
