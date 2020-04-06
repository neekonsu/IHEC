package main

import (
	"flag"
	"fmt"

	ihec "github.com/neekonsu/IHEC"
)

func main() {
	jsonPath := flag.String("jsonPath", "./res", "Path to directory storing JSON metadata files")
	selection := ihec.PopulateFiles(*jsonPath)
	selection.PopulateAccessions()
	for _, item := range selection.Accessions {
		fmt.Println(item)
	}
}
