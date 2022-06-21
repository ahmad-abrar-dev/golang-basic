package helper

//access modifier is mean , to set an access of a function or variable into a public or private of a PACKAGE
// to use public function/variable make sure named these stuff with Uppercase in the first letter
// to use private function/variable make sure named these stuff with lowercase in the first letter
var version = "1.0.0"      //private var
var Application = "golang" //public var

// *** these function implemented into import.go

//private function
func sayGoodBy(name string) string {
	return "Bye " + name
}

//public function
func SayHello(name string) string {
	return "Hello" + name
}
