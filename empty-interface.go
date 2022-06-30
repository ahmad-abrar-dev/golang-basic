package main

import "fmt"

// the empty interface is an interface that hasn't implement any struct
// it can return or implement every data type
// so basically implementing the empty interface is when we need all of the data type can be return
// in other word the empty interface do not even care whatever the data type is

// declarating an empty interface
// type Whatever interface {}
// type Whatever interface {
//
// }

// func Ups(i int, whatever interface{}) interface{} {
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	//but to implement the empty interface we should describe the data type
	var data interface{} = Ups(1)
	fmt.Println(data)
}
