package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5}
	for id := range ids {
		fmt.Printf("This is the number %d\n", ids[id])
	}

	sum := 0
	for id := range ids {
		sum += ids[id]
	}
	fmt.Printf("The sum is %d\n", sum)
}
