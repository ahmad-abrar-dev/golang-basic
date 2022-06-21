package main

import (
	"fmt"
)

func random() interface{} {
	return "OK"
}

func main() {
	// result := random()
	// fmt.Println(result)

	// result.(string) will convert the interface into a string, but
	// we need to make sure the return of the function is a string (same as what we converted) or the app will panic (shut down)
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int) //panic
	// fmt.Println(resultInt)

	//use switch to make it more safety
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
}
