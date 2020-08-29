package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

type lecture struct {
	name       string
	instructor person
	duration   int
}

func main() {
	lectures := []lecture{
		lecture{"golang", person{"vicky", "kumar"}, 350},
		lecture{"oncurrency", person{"nanthan", "nanetha"}, 300},
		lecture{"Gin", person{"Ipavum nantha", "marupadium nanetha"}, 250},
	}

	for _, lecture := range lectures {
		name := lecture.name
		instructor := fmt.Sprintf("%s %s", lecture.instructor.firstname, lecture.instructor.lastname)
		duration := lecture.duration

		fmt.Printf("Lecture: '%s', Author: '%s', Duration: %d secs\n", name, instructor, duration)
	}
}
