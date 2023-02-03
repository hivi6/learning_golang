package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string
	// assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	fmt.Println(fruitArr, fruitArr[0])

	// Declare and assign
	anotherFruitArr := [2]string{"Apple", "Orange"}
	fmt.Println(anotherFruitArr)

	// Slices
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}
	fmt.Println(fruitSlice, len(fruitSlice))
	fmt.Println(fruitSlice[1:3]) // [1, 3)
}
