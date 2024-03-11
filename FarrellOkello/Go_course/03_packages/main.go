package main

import (
	"fmt"
	"math"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func greeting(name string) string {
	return name
}

func getProduct(a int, b int) int {
	return a * b
}

func main() {
	fmt.Println("Hello ", greeting("Farrel Manigga"))
	fmt.Println("The product of A and B is :", getProduct(12, 89))
	fmt.Println("The reverse of Hello is :", Reverse("Hello"))
	fmt.Println(math.Floor(90.7))
	fmt.Println(math.Ceil(90.7))
	fmt.Println(math.Sqrt(16))
}
