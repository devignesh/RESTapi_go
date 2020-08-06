package main

import (
	"fmt"
	"strconv"
)

func strtoint() {
	fmt.Println("String to integer conversion")

	var value int

	value, err := strconv.Atoi("3232")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(value + 1)
}

func inttostr() {
	fmt.Println("integer to string conversion")
}

func main() {
	fmt.Println("String type conversion")

	strtoint()

}
