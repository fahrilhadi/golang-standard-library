package main

import (
	"container/list"
	"fmt"
)

func main()  {
	data := list.New()

	data.PushBack("Fahril")
	data.PushBack("Hadi")

	head := data.Front()
	fmt.Println(head.Value) // Fahril

	next := head.Next()
	fmt.Println(next.Value) // Hadi

	// menggunakan perulangan

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}