package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

// this can also be written as
// func getSum(num1, num2 int) int {
func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("bob"))
	fmt.Println(getSum(1, 2))
}
