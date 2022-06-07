package main

import "fmt"

func main() {
	var name string

	name = "zhun weky"
	fmt.Println(name)

	name = "ahmad abrar"
	fmt.Println(name)

	var testVar = "testVariable"
	fmt.Println(testVar)

	//if int var not declarated it will automatically int32/int64 depends on your bit OS
	var age int8 = 30
	fmt.Println(age)

	// use := it will not require the var statement
	country := "Indonesia"
	fmt.Println(country)

	//multiple variable declaration
	var (
		firstName = "zhun"
		lastName  = "weky"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
