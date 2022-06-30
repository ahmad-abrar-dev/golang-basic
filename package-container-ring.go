package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

//documentation : https://golang.org/pkg/container/ring/

//like the container list, but in container ring there is no end of data, it will return back to the first data
func main() {
	// var data *ring.Ring = ring.New(5)
	data := ring.New(5)
	data.Value = "zhun"

	// var data2 = data.Next()
	// data2.Value = "test"
	// fmt.Println(*data2)

	for i := 0; i < data.Len(); i++ {
		// data.Value = "Value " + strconv.FormatInt(int64(i), 10)
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	// use do loop
	// the Do() loop is a loop just like a foreach loop
	// the Do() loop just can use on container ring datas
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
