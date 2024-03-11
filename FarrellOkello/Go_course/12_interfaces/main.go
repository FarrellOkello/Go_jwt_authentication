package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	x, y, raduis float64
}

type Rectangle struct {
	width, length float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.raduis * c.raduis
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}

func getARea(s Shape) float64 {
	return s.Area()
}

func main() {
	circle := Circle{x: 0, y: 0, raduis: 20}
	rectangle := Rectangle{width: 5, length: 10}
	fmt.Printf("The area of the circle %f\n", getARea(circle))
	fmt.Printf("The area of the retangle %f\n", getARea(rectangle))

}
