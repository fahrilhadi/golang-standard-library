package main

import (
	"encoding/csv"
	"os"
)

func main()  {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"fahril", "hadi"})
	_ = writer.Write([]string{"budi", "pratama"})
	_ = writer.Write([]string{"joko", "diah"})

	writer.Flush()
}