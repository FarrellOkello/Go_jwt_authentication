package main

import "fmt"

func main() {
	// /arrays have a fixed number
	var array = [2]string{"Apple", "Orange"}
	fmt.Println(array)
	fmt.Println(array[1])
	arraySlice := []string{"Jago", "JJ", "Professor"}
	fmt.Println(arraySlice[1:3])
}
