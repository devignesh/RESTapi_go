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

	fmt.Println(value)
}

func inttostr() {
	fmt.Println("integer to string conversion")

	var value1 string

	value1 = strconv.Itoa(3232)

	fmt.Println(value1)
}

func main() {
	fmt.Println("String type conversion")

	strtoint()
	inttostr()

}
