package main

import "fmt"

func sayHello() {
	fmt.Println("Hello!")
}

// function using parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello!", firstName, lastName)
}

// function with return data
func getHello(name string) string {
	return "Hello " + name
}

// function with return multiple data
func greeting() (string, string) {
	return "ahmad", "abrar"
}

//function named return values
//it can store return with variable existed in the function
func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "ahmad"
	middleName = "abrar"
	lastName = "r"
	// return firstName, middleName, lastName
	return
}

//variadic function
//these variadic argument variable required in the last argument of a function as a final argument
//it can't be used on the firt variable if u'r use 2 variable as variadic argument
//the variadic
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//function as value
func getGoodBye(name string) string {
	return "Good Bye " + name
}

//function as parameter
// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello", nameFiltered)
// }
//we even can use type declaration and the result same as above
// the type declaration describe about the structure of a function
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

//anonymous function
// function who called without name
// generally these function called on parameter where the function required a function to implement as parameter
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you're blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// recursive function
// this is the function who called them self
// it's looks like a loop function
// it's alson required a stopper on the function and need to make sure the loop to stop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}

}

func main() {
	sayHello()
	sayHelloTo("zhun", "weky")

	//call a function with return data
	result := getHello("abrar")
	fmt.Println(result)

	//call function with multiple return
	firstName, lastName := greeting()
	fmt.Println(firstName)
	fmt.Println(lastName)

	//the "_" (underscore) used for ignoring value when we don't need the variable
	_, ln := greeting()
	fmt.Println(ln)

	//function named return values
	fn, mn, ln := getFullName()
	fmt.Println(fn)
	fmt.Println(mn)
	fmt.Println(ln)

	// implement variadic function
	total := sumAll(10, 10, 10, 10, 10, 10)
	fmt.Println(total)
	// we also can use slice as parameter on variadic function
	// pas the slice variable with triple dot "..."
	slice := []int{20, 20, 20, 20}
	fmt.Println(sumAll(slice...))

	//implement function as value
	goodBye := getGoodBye
	fmt.Println(goodBye("zhun"))

	//implement function as parameter
	sayHelloWithFilter("Eko", spamFilter)
	// or
	filterr := spamFilter
	sayHelloWithFilter("anjing", filterr)

	//implement anonymous function
	Blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("zhun", Blacklist)
	registerUser("admin", Blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})

	//implement recruisive function
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}
