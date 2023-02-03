package main

import "fmt"

var globalVar1 = "This is a global variable"

// globalVar2 := "This is a shorthand global variable" // This is an error

func main() {
	// MAIN TYPES
	// string
	// bool
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// printing global variable
	fmt.Println(globalVar1)

	// create variables using var keyword
	var name string = "Hello"
	fmt.Println(name)

	// create variables without variable type by using inference
	var name2 = "Hello2"
	fmt.Println(name2)

	// in go if you create a variable but don't use it then it's an error
	// below line is an error
	// var unusedVar int = 17

	// print the type of a given variable
	fmt.Printf("%T\n", name)

	// changing value
	var a = true
	a = false
	fmt.Println(a)

	// const values cannot be change
	const b = true
	// b = false // this line is an error
	fmt.Println(b)

	// using the shorthand method for variable assignment
	shorthandVar := "This is a shortHand Var"
	fmt.Println(shorthandVar)

	// multiple variable
	email, pass := "example@domain.com", "password1"
	fmt.Println(email, pass)
}
