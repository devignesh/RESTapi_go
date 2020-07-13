package main

import (
	"fmt"
	"syscall/js"
)

func add(i []js.Value) {
	js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
	println(js.ValueOf(i[0].Int() + i[1].Int()).String())
}

func substract(i []js.Value) {
	js.Global().Set("output", js.ValueOf(i[0].Int()-i[1].Int()))
	println(js.ValueOf(i[0].Int() - i[1].Int()).String())
}

func multiply(i []js.Value) {
	js.Global().Set("output", js.ValueOf(i[0].Int()*i[1].Int()))
	println(js.ValueOf(i[0].Int() * i[1].Int()).String())
}

func division(i []js.Value) {
	js.Global().Set("output", js.ValueOf(i[0].Int()/i[1].Int()))
	println(js.ValueOf(i[0].Int() / i[1].Int()).String())
}

func registerallbacks() {
	js.Global().Set("add", js.NewCallback(add))
	js.Global().Set("subtract", js.NewCallback(substract))
	js.Global().Set("multiply", js.NewCallback(multiply))
	js.Global().Set("division", js.NewCallback(division))
}

func main() {
	fmt.Println("hii hello")

	c := make(chan struct{}, 0)

	registerallbacks()
}
