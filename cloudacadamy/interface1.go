package main

import "fmt"

type device interface {
	turnOn() string
}

type mobile struct {
	name  string
	model string
}

type laptop struct {
	name  string
	model string
}

func (phone mobile) turnOn() string {
	return "Mobile starting up..."
}

func (lap laptop) turnOn() string {
	return "Laptop starting up..."
}

func main() {
	dev1 := mobile{"oneplus", "8 Pro"}
	dev2 := laptop{"mackbook pro", "1.79l"}

	devices := []device{dev1, dev2}

	for _, dev := range devices {
		fmt.Println(dev.turnOn())
	}
}
