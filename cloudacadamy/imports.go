package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

import "strconv"
import m "math"

func main() {
	name := "vicky"

	fmt.Println(strings.ToUpper(name))
	fmt.Println(uuid.New())
	f, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f)
	fmt.Println(m.Round(f))
}
