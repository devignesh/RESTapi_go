package main

import "fmt"

func doSomething(msg string) {
	fmt.Println(msg)
}

func system() int {
	fmt.Println("system started func...")

	defer doSomething("cleanup")
	defer doSomething("stop")

	fmt.Println("system finished!")

	return 1
}

func main() {
	data := system()
	fmt.Println(data)
}
