package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main()  {
	csvString := "fahril, hadi\n" +
	"budi, pratama\n" +
	"joko, morro"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}