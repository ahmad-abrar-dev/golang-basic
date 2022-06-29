package main

import (
	"container/list"
	"fmt"
)

// documentation : https://golang.org/pkg/container/list/

func main() {
	//linked list example (data structure)
	//it's just like an array or slice but no index
	//we just can access linked list data in the begining of data or the first and the last of the datas
	// use Next() to access the next data or the Prev() to access the previous of data
	data := list.New()
	// the pushBack() used to add data as last of data
	// the pushFront() used to add data as the first of data
	// the Len() to count the length of the data list
	// the Remove() to remove data
	data.PushBack("data1")
	data.PushBack("data2")
	data.PushBack("data3")
	data.PushBack("data4")
	data.PushBack("data5")
	data.PushBack("data6")
	data.PushFront("data0")

	fmt.Println("---------------")
	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)
	fmt.Println(data.Front().Next().Value)
	fmt.Println(data.Back().Prev().Prev().Value)
	fmt.Println("---------------")

	//access linked list from front to back
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	//access linked list from back to front
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

}
