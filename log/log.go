package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"time"
)

func AddRandomNumbers() {
	firstnumber := rand.Intn(100)
	secondnumber := rand.Intn(100)

	time.Sleep(2 * time.Second)

	result := firstnumber + secondnumber

	fmt.Printf("result of addition %d\n", result)

}
