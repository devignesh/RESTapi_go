package main

import (
	"fmt"
	"sync"
	"time"
)

func waitgp(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("finish goroute")
	wg.Done()

}

func main() {
	fmt.Println("waitgroup tutorial")
	var wg sync.WaitGroup
	wg.Add(1)
	go waitgp(&wg)
	wg.Wait()
	fmt.Println("finish execution of waitgroup")

}
