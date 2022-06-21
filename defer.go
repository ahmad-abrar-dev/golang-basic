package main

import "fmt"

// the defer will be execute after the code in the function itself has been executed
// no matter if the function error or no

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runAppllication(value int) {
	defer logging()
	fmt.Println("Run application")
}

func main() {
	runAppllication(10)
}
