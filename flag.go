package main

import (
	"flag"
	"fmt"
)

// the flag package usage basically to adding a flag on cli e.g: go run flag.go -host:localhost
// example usage
// go run flag.go -host=localhost -user=toor -password=root123
func main() {
	// flag.<datatype>("<specified/indicator/flag name>", "<default value>", "<message described when err")
	// it will return a string address for this string function
	// flag.String("host", "localhost", "Put your host")
	var host *string = flag.String("host", "localhost", "Put your host")
	var user *string = flag.String("user", "root", "Put your root")
	var password *string = flag.String("password", "root", "Put your password")
	var number *int = flag.Int("number", 100, "Put your number occured")

	// to do parsing process we need to define it
	// it will call default value when we're not call this method
	// so when we set the value in cli those flag will not work
	flag.Parse()

	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Password : ", *password)
	fmt.Println("Password : ", *number)

	//to get the pointer value use *
	//e.g : *host

}
