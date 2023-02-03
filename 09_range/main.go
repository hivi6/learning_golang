package main

import "fmt"

func main() {
	ids := []int{123, 12, 3123, 14, 3, 13}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// using range with map
	emails := map[string]string{"Ken": "ken@domain.com", "Bob": "bob@domain.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
	for k := range emails {
		fmt.Printf("Name: %s\n", k)
	}
}
