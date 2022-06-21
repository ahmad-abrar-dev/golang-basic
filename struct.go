package main

import "fmt"

// type Customer struct {
// 	Name    string
// 	Address string
// 	Age     int
// }

// make it short as like this
type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var client Customer
	client.Name = "zhun"
	client.Address = "tangerang"
	client.Age = 27

	fmt.Println(client)
	fmt.Println(client.Name)
	fmt.Println(client.Address)
	fmt.Println(client.Age)

	client2 := Customer{
		Name:    "joko",
		Address: "Indonesia",
		Age:     24,
	}
	fmt.Println(client2)

	// it shorter but we need to make sure the structure doesn't change or it will be error
	client3 := Customer{"Budi", "Indonesia", 30}
	fmt.Println(client3)
}
