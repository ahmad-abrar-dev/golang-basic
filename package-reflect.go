package main

import (
	"fmt"
	"reflect"
)

// docs : https://golang.org/pkg/reflect
// so basically the package reflect is a package used to read the data struct

type Sample struct {
	Name string `required:"true" max:"10"` //tags of struct
	Age  int    `required:"true" min:2`
}

type AnotherSample struct {
	Name string
	Desc string
}

//validation
func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			// return reflect.ValueOf(data).Field(i).Interface() != ""
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"zhun", 12}

	// var sampleType reflect.Type = reflect.TypeOf(sample)
	sampleType := reflect.TypeOf(sample)

	// structField := sampleType.Field(0)

	// fmt.Println(sampleType.NumField())
	// fmt.Println(sampleType.Field(1).Name)
	// fmt.Println(structField.Name)
	required := sampleType.Field(0).Tag.Get("required") // get the value of tag in field 0 -> Name
	otherRequired := sampleType.Field(0).Tag.Get("max")

	fmt.Println(required)
	fmt.Println(otherRequired)

	//try validation
	fmt.Println("==============================")
	// sample.Name = ""
	// sample.Age = 0
	fmt.Println(isValid(sample))
	fmt.Println("==============================")

	//in AnotherSample there's no tag, so it  will return true in isValid function
	sampleAgain := AnotherSample{"", ""}
	fmt.Println(isValid(sampleAgain))

}
