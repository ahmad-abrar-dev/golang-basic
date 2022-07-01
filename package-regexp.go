package main

import (
	"fmt"
	"regexp"
)

// documentations : https://golang.org/pkg/regexp/ , https://github.com/google/re2/wiki/syntax
// the regexp stand for regular expresion
// regexp it's self is an utility in Go to find a regular expression
// regexp in go used C language that who created by Google called RE2
func main() {
	// regexp.MustCompile(string) // create a regexp
	// regexp.MatchString(string) bool // checking the regexp matched withe the string
	// regexp.FindAllString(string, max) bool // finding match string with the maximum result
	var regex *regexp.Regexp = regexp.MustCompile(`z([a-z])w`) //([a-z]) determinate each word

	fmt.Println(regex.MatchString("zaw"))
	search := regex.FindAllString("zun zaw ziw zAr", 2) // it will return slice of string
	//the second parameter will find how much to show, and -1 to find all of it
	fmt.Println(search)
}
