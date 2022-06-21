package main

import "fmt"

type Man struct {
	Name string
}

//a method
func (man *Man) Married() {
	// use (*) sign as the example above to indicate the method use pointer
	// so the variable who use these function will get change.
	man.Name = "Mr. " + man.Name
}

func main() {
	bro := Man{"zhun"}
	bro.Married()

	fmt.Println(bro.Name)
}
