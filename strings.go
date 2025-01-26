package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(strings.Contains("Fahril Hadi", "Fahril"))
	fmt.Println(strings.Split("Fahril Hadi", " "))
	fmt.Println(strings.ToLower("Fahril Hadi"))
	fmt.Println(strings.ToUpper("Fahril Hadi"))
	fmt.Println(strings.Trim("     Fahril Hadi     ", " "))
	fmt.Println(strings.ReplaceAll("Fahril Hadi Fahril Hadi", "Fahril", "Fadli"))
}