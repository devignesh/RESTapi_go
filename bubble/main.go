package main

import (
	"fmt"
)

var sorte [10]int = [10]int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}

func bubblesort(input [10]int) {

	n := 10
	swapped := true

	for swapped {
		swapped = false

		for i := 1; i < n; i++ {

			if input[i-1] > input[i] {
				fmt.Println("swappping element")

				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
	fmt.Println(input)
}

func main() {
	fmt.Println("Bubble sort")

	bubblesort(sorte)
}
