package main

import (
	"fmt"
)

func bubble(numbers []int) {
	var n int = len(numbers)

	var firstindex int = 0
	var secondindex int = 1

	for secondindex < n {
		var firstnumber int = numbers[firstindex]
		var secondnumber int = numbers[secondindex]

		// fmt.Println("comparision", firstnumber, secondnumber)

		if firstnumber > secondnumber {
			numbers[firstindex] = secondnumber
			numbers[secondindex] = firstnumber

		}

		firstindex++
		secondindex++
	}
}

func sortloop(numbers []int) {
	var n int = len(numbers)

	for i := 0; i < n; i++ {
		bubble(numbers)
		fmt.Println("vicky", numbers)
	}
}

func main() {
	var numbers []int = []int{5, 8, 1, 4, 3, 2, 7, 0}
	fmt.Println("numbeers", numbers)
	bubble(numbers)
	sortloop(numbers)
	fmt.Println("sort", numbers)

}
