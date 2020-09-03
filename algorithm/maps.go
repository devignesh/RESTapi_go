package main

import "fmt"

func main() {

	emails := map[string]string{"vicky": "vicky@gmail.com", "kumar": "kumar@gmail.com"}

	emails["ducky"] = "ducky@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["vicky"])

	// Delete from map
	delete(emails, "ducky")
	fmt.Println(emails)
}
