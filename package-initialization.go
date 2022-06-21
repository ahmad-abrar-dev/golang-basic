package main

import (
	"fmt"
	"golang-basic/database"
	// _ "golang-basic/database"
	// blank identifier is a way to call something but not in using, it's also can implement to imported package
	// espectically when we just need to execute the init function on the package
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)

}
