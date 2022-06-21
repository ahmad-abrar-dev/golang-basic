package main

import "fmt"

type HasName interface {
	// describe the contract param and or the return type data
	// the contract described here should be use by the object who use the interface (so we need to create the function)
	// the interface intended to make sure a struct or a class to implement the contract described in interface
	GetName() string
	GetAddress() string
}

func SayHello(hasName HasName) { // the HasName here described the interface
	fmt.Println("Hello", hasName.GetName())
	fmt.Println("your address:", hasName.GetAddress())
}

type Person struct {
	Name    string
	Address string
}

func (person Person) GetName() string {
	return person.Name
}

func (person Person) GetAddress() string {
	return person.Address
}

//sample 2
type Animal struct {
	Name    string
	Address string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (animal Animal) GetAddress() string {
	return animal.Address
}

func main() {
	var man Person
	man.Name = "zhun"
	man.Address = "tangerang"

	SayHello(man)
	animal := Animal{
		Name:    "cat",
		Address: "Home",
	}
	SayHello(animal)
}
