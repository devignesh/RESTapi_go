package main

import "fmt"

func main() {
	nums := []int{4, 22, 2, 1, 333, 343}

	// Loop through ids
	for i, num := range nums {
		fmt.Printf("%d - Num: %d\n", i, num)
	}

	// Not using index
	for _, num := range nums {
		fmt.Printf("Num: %d\n", num)
	}

	// Add ids together
	sum := 0
	for _, id := range nums {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"vicky": "vicky@gmail.com", "kumar": "kumar@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
