package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

/*
Above struct can also be written as
type Person struct {
	firstName, lastName, city, gender string
	age int
}
*/

// greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and i'm " + strconv.Itoa(p.age) + " years old."
}

// has birthday (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// get married (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "F" {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init Person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "F", age: 25}
	fmt.Println(person1)
	fmt.Println(person1.greet())

	person1.hasBirthday()
	fmt.Println(person1.greet())

	person1.getMarried("Williams")
	fmt.Println(person1.greet())
}
