package main

import (
	"flag"

	ihec "github.com/neekonsu/IHEC"
)

func main() {
	jsonPath := flag.String("jsonPath", "./res", "Path to directory storing JSON metadata files")
	selection := ihec.PopulateFiles(*jsonPath)
	selection.PopulateAccessions()
	selection.PrintAccessions()
	selection.PrintLeanContext()
}
