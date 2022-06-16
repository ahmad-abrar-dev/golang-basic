package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "zhun weky",
		"address": "tangerang",
	}
	person["title"] = "proggrammer"

	fmt.Println(person)
	fmt.Println(person["name"])

	//remove map
	delete(person, "address")
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "pzn book"
	book["author"] = "eko khanedy"

	fmt.Println(book)

}
