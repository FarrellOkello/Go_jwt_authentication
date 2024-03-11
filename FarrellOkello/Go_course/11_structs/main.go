package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	gender    string
	city      string
	age       int
}

func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + ", I come from " + p.city + " i am " + strconv.Itoa(p.age) + " years old.\n"
}

func (p *Person) getMarried(lastName string) {
	if p.gender == "f" {
		p.lastName = lastName
	}
}

func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	person1 := Person{firstName: "Farrel", lastName: "OKello", gender: "m", city: "New York", age: 34}
	person1.hasBirthday()
	person1.getMarried("Jaspher")
	fmt.Printf(person1.greet())
	fmt.Println(person1)
}
