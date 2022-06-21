package main

import "fmt"

func endApp() {
	message := recover() //recover use to catch error message from panic function without exit the app
	if message != nil {
		fmt.Println("error with message :", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(eror bool) {
	defer endApp()
	if eror {
		panic("Aplikasi Error") // the panic function use for exit app with error message and stop executing
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("yes the app still running")
}
