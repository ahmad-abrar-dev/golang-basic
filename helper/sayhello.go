package helper

// the rule in creating a package is in one package can't have a same named function
// to use the package we need to use "import"

import "fmt"

func TestSayHello(name string) {
	fmt.Println("Hello", name)
}
