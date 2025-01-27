package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main()  {
	users := []User {
		{"Fahril", 25},
		{"Hadi", 24},
		{"Fadli", 23},
		{"Ramdani", 22},
		{"Abu", 21},
		{"Hanif", 20},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}