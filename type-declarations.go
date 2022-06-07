package main

import "fmt"

func main() {
	type NoKTP string
	type handsome bool
	var myKTP NoKTP = "21302830123"
	var male handsome = true

	fmt.Println(myKTP)
	fmt.Println(male)
}
