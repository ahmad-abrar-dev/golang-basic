package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// this way (pointer usage) required to change a middle/big size struct
// to decrease the memory usage
func ChangeAddressToIndonesia(address *Address) {
	// add (*) sign as above example into the argument datatype to indicate the pointer
	// so basically we can made a change into the variable wherever that variable is (outside this function),
	// it's just like a global scoped variable
	address.Country = "Indonesia"

}

func main() {
	address := Address{"Pekanbaru", "Riau", ""}
	// add (&) sign as below example to pass the pointer value bcz in the func argument required pointer
	ChangeAddressToIndonesia(&address)

	fmt.Println(address)
}
