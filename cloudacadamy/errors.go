package main

import (
	"errors"
	"fmt"
	"math"
)

func circleArea(radius float32) (float32, error) {
	if radius < 0 {
		return 0, errors.New("radius should be positive value")
	} else {
		return math.Pi * radius * radius, nil
	}
}

func main() {
	area1, _ := circleArea(3)
	fmt.Println(area1)

	if area2, err := circleArea(-3); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area2)
	}
}
