package main

import (
	"fmt"
	"strings"
)

type device interface {
	turnOn() string
	update(version float32)
}

type mobile struct {
	name    string
	model   string
	version float32
}

type laptop struct {
	name    string
	model   string
	version float32
}

func (phone mobile) turnOn() string {
	return "Mobile starting up..."
}

func (lap laptop) turnOn() string {
	return "Laptop starting up..."
}

func (ph *mobile) update(version float32) {
	ph.version = version
}

func (laptops *laptop) update(version float32) {
	laptops.version = version
}

func main() {
	dev1 := mobile{"oneplus", "8 Pro", 1.0}
	dev2 := laptop{"mackbook pro", "1.79l", 8.1}

	devices := []device{&dev1, &dev2}
	fmt.Println(devices)

	for _, dev := range devices {
		if strings.Contains(dev.turnOn(), "Oneplus") {
			dev.update(1.1)
		} else if strings.Contains(dev.turnOn(), "mackbook") {
			dev.update(8.12)
		}
		fmt.Println(dev.turnOn())
	}

	fmt.Println(dev1)
	fmt.Println(dev2)
}
