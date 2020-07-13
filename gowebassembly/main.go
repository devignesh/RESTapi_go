package main

import (
	"fmt"
	"syscall/js"
)

func add(i []js.Value) {
	js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
	println(js.ValueOf(i[0].Int() + i[1].Int()).String())
}

func registerallbacks() {
	js.Global().Set("add", js.NewCallback(add))
}

func main() {
	fmt.Println("hii hello")

	c := make(chan struct{}, 0)

	registerallbacks()
}
