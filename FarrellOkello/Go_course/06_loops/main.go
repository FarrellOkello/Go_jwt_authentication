package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("Number %d\n", i)
	}

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("Fizz")
		} else if i%3 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Printf("Number %d\n", i)
		}
	}
}
