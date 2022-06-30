package main

import (
	"fmt"
	"sort"
)

//docs : https://golang.org/pkg/sort/

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	users := []User{
		{"Zhun", 30},
		{"Weky", 20},
		{"Diki", 50},
		{"Bintang", 10},
	}

	//call the sort method
	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
