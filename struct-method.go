package main

import "fmt"

type Customer struct {
	Name, Address string
	age           int
}

//if we create a function that who specific of a struct it will be better to make like this
func sayHiOld(customer Customer, name string) {
	fmt.Println("Hello ", name, "My Name is ", customer.Name)
}

// to implement the following code above is like sayHi(customer, name) like a normal methos
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello ", name, "My Name is ", customer.Name)
}

//the variable doesn't require same as like the struct name
func (x Customer) sayOops(name string) {
	fmt.Println("Oops", name, "you're not", x.Name)
}

func main() {
	var client Customer
	client.Name = "zhun"
	client.Address = "tangerang"
	client.age = 30

	client.sayHi("abrar")
	client.sayOops("abrar")
}
