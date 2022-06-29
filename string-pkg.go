package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("zhun weky", "weky"))                 // it will return boolean as true
	fmt.Println(strings.Split("zhun weky", " "))                       // split string to an array of slice (slice of string)
	fmt.Println(strings.ToLower("Zhun Weky"))                          // lowercase
	fmt.Println(strings.ToUpper("Zhun Weky"))                          // uppercase
	fmt.Println(strings.ToTitle("zhun weky"))                          // same as uppercase
	fmt.Println(strings.Trim("_____Zhun_Weky______", "_"))             // remove corner character
	fmt.Println(strings.ReplaceAll("zhun weky zhun", "zhun", "abrar")) // replace all character contained
}
