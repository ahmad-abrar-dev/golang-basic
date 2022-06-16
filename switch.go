package main

import "fmt"

func main() {
	name := "zhun weky"

	switch name {
	case "zhun":
		fmt.Println("hello Zhun")
	case "weky":
		fmt.Println("hello weky")
	default:
		fmt.Println("who are you?")
	}

	//switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("it's to long")
	case false:
		fmt.Println("it's okay")
	}

	//switch without condition
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("the name is to long")
	case length > 5:
		fmt.Println("the name is okay")
	default:
		fmt.Println("the name is to short")
	}

}
