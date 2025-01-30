package main

import (
	"fmt"
	"regexp"
)

func main()  {
	regex := regexp.MustCompile(`a([a-z])i`)

	fmt.Println(regex.MatchString("adi"))
	fmt.Println(regex.MatchString("ari"))
	fmt.Println(regex.MatchString("aRi"))

	fmt.Println(regex.FindAllString("adi aji ade ari aci aKi",10))
}