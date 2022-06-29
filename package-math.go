package main

import (
	"fmt"
	"math"
)

// documentation : https://golang.org/pkg/math/

func main() {
	fmt.Println(math.Round(12.51)) // round up > 5
	fmt.Println(math.Floor(12.7))  // round down
	fmt.Println(math.Ceil(12.1))   // round up
	fmt.Println(math.Max(10, 20))  // get max value both of these parameter
	fmt.Println(math.Min(10, 20))  // get min value both of these parameter
}
