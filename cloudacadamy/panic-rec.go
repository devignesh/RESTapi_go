package main

import "fmt"

func panicfunc() int {
	fmt.Println("system started...")

	defer func(msg string) {
		if r := recover(); r != nil {
			fmt.Println("recovered")
		}
		fmt.Println(msg)
	}("Vicky")

	var data []int
	var x = data[0]
	x++

	fmt.Println("system finished!")

	return 1
}

func main() {
	data := panicfunc()
	fmt.Println(data)

	panic("die!")
}
