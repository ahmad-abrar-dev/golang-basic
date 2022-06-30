package main

import (
	"fmt"
	"strconv"
)

//strconv stand for string conversion
// completed documentation : https://golang.org/pkg/strconv/

//almost all stconv package return error as 2nd return
func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	//base size
	/*
		decimal = 10
		binary = 2
		octal = 8
		hexa = 16
	*/
	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	//change number format like int to the base what we want
	value := strconv.FormatInt(1000000, 2)
	fmt.Println(value)

	//convert int to str
	intToStr := strconv.Itoa(20)

	//convert str to int
	strToInt, _ := strconv.Atoi("20")

	fmt.Println(intToStr)
	fmt.Println(strToInt)
}
