package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "zhun"
	names[1] = "weky"
	names[2] = "ahmad"

	fmt.Println(names)
	fmt.Println(names[1])

	////////////// make an array in one line, if the value not declarated it will be 0
	var values = [3]int{
		90,
		91,
	}

	fmt.Println(values)

	///////////// declarating an array without describing the length
	var valuesWithoutRange = [...]int{
		90,
		91,
		92,
		93,
		94,
	}

	fmt.Println(valuesWithoutRange)


	//the len function used to know the range of an array --note:not the values included
	fmt.Println(len(valuesWithoutRange))\
}
