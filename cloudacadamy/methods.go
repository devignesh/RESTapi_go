package main

import "fmt"

type driving struct {
	firstname string
	surname   string
	age       int
}

func (d *driving) namecon() string {
	return d.firstname + " " + d.surname
}

func (d *driving) candrive() bool {
	if d.age >= 20 {
		return true
	} else {
		return false
	}
}

func (d *driving) ageupdate(newAge int) {
	d.age = newAge
}

func main() {
	person1 := driving{"vicky", "kumar", 35}
	person2 := driving{"vignesh", "kumar", 18}

	fmt.Printf("%s can drive: %t\n", person1.namecon(), person1.candrive())
	fmt.Printf("%s can drive: %t\n", person2.namecon(), person2.candrive())

	person2.ageupdate(person2.age + 1)
	fmt.Println(person2.age)

	fmt.Printf("%s can drive: %t\n", person2.namecon(), person2.candrive())
}
