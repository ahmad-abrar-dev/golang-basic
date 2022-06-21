package main

import "fmt"

// the nil value same as null in other proggramming language
// the nil value commonly use when we need to describe the value as empty
// where we know the return value should implement like we have described in a function
// so we can set a nil value to implement the null return
// the nil value can set on interface, function, map, slice, pointer and channel
// other else will return default value like int = 0 , string = ""
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// var person map[string]string = nil
	person := NewMap("Eko")
	fmt.Println(person)
}
