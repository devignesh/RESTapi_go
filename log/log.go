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

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("Trace file is not created %v\n", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("faile to close %v\n", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start: %d\n", err)
	}
	defer trace.Stop()

	AddRandomNumbers()
}
