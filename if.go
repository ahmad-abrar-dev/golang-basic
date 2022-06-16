package main

import "fmt"

func main() {

	//basic usage if statement
	name := "budisa"

	if name == "zhun" {
		fmt.Println("hi " + name)
	} else if name == "budi" {
		fmt.Println("you're " + name)
	} else {
		fmt.Println("we don't know who are you")
	}

	//short statement
	if length := len(name); length > 5 {
		fmt.Println("your name is to long")
	} else {
		fmt.Println("the name is okay")

	}

}
