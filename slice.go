package main

import "fmt"

func main() {
	// slice is a method to slicing the array, it will get the impact into existing array's data.
	//creating new slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "zhun"
	newSlice[1] = "weky"

	//we can use an apeend method to the slice

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	///////////// it's a slice
	var valuesWithoutRange = []int{
		90,
		91,
		92,
		93,
	}

	fmt.Println(valuesWithoutRange)

}
