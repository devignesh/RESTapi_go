package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pause() {
	n := rand.Intn(3)
	time.Sleep(time.Duration(n) * time.Second)
}

func doSomething(msg string) {
	pause()
	fmt.Println(msg)
}

func main() {
	rand.Seed(time.Now().Unix())

	doSomething("sync1")

	go doSomething("async1")
	go doSomething("async2")
	go doSomething("async3")

	doSomething("sync2")

	time.Sleep(time.Second * 10)
}
