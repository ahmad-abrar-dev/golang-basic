package main

import "fmt"

func main() {
	counter := 1

	// it's just like a while loop in other proggramming language
	// in go-lang there's just only for loop available
	for counter <= 10 {
		fmt.Println("Loop -", counter)
		counter++
	}

	//but in golang we can do some statement (init statement / post statement)
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("2nd Loop -", counter)
	}

	//a normal for loop in slice
	slice := []string{"zhun", "weky", "ahmad", "abrar"}

	// fmt.Println("get single slice", slice[1])
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//for range
	// if not use the index or value just use "_" (underscore)
	for index, value := range slice {
		fmt.Println("index -", index, "=", value)
	}
}
