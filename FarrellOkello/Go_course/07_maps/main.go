package main

import "fmt"

func main() {
	// // maps
	// emails := make(map[string]string)
	// emails["Bob"] = "bob@gmail.com"
	// emails["Farrel"] = "farrel@gmail.com"
	// emails["Farrel"] = "farrel@gmail.com"

	// shothand way to create and initialize a map
	emails := map[string]string{"Bob": "bob@gmmail.com", "Farrell": "farrell@outlook.com"}

	fmt.Println(emails["Bob"])
	fmt.Println(emails)
}
