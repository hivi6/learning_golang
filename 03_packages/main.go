package main

import (
	"fmt"
	"math"

	"github.com/hivi6/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(1.4))
	fmt.Println(math.Ceil(1.4))
	fmt.Println(math.Sqrt(1.4))
	fmt.Println(strutil.Reverse("hello"))
}
