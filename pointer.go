package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// use & indicate "go to the address"
	// use * indicate "every var with the same address go here"
	address1 := Address{"Subang", "JaBar", "Indonesia"}
	address2 := &address1 // create var2 and point to the address 1

	// those variables following above is equal as the code below
	// var address1 Address = Address{"Subang", "JaBar", "Indonesia"}
	// var address2 *Address = &address1

	//it will store into the same address
	// address2.City = "Bandung"

	//it will create new address
	// address2 = &Address{"Jakarta", "DKI", "Indonesia"}

	//whoever use the same addres will follow the *var address
	*address2 = Address{"Jakarta", "DKI", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)
}
