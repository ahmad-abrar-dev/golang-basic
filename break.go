package main

import "fmt"

func main() {

	//break use for stoping the loop
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke ", i)
	}

	// continue use for do next step the loop
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke ", i)
	}

	// break out nested loop
out:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			println("i:", i, ", j:", j)
			if i+j == 10 {
				break out
			}
		}
	}

}
