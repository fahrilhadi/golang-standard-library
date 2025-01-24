package main

import "fmt"

func main()  {
	firstName := "Fahril"
	lastName := "Hadi"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}