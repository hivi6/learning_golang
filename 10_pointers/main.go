package main

import "fmt"

func main() {
	a := "anc"
	a_address := &a

	fmt.Println(a, a_address)
	fmt.Printf("%T\n", a_address)

	fmt.Println(*a_address)
	*a_address = "xyz"
	fmt.Println(a)
}
