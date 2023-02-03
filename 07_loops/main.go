package main

import "fmt"

func main() {
	// Long method (like while loop in c)
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// short method
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		ans := ""
		if i%3 == 0 {
			ans += "Fizz"
		}
		if i%5 == 0 {
			ans += "Buzz"
		}
		if ans == "" {
			fmt.Println(i)
		} else {
			fmt.Println(ans)
		}
	}
}
