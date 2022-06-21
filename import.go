package main

import (
	"fmt"
	"golang-basic/helper"
)

// *** these function implemented from helper package
//note: in one app there's only 1 main function
func main() {
	testImport := helper.SayHello("Zhun")
	fmt.Println(testImport)

	// testImport2 := helper.sayGoodBye("Zhun") // it will return error bcz the sayGoodBye function is a private
	// fmt.Println(testImport2)

	//trying access variable
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) //error

}
