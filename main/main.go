package main

import (
	"flag"

	ihec "github.com/neekonsu/IHEC"
)

func main() {
	jsonPath := flag.String("jsonPath", "./res", "Path to directory storing JSON metadata files")
	outputAccessionsPath := flag.String("outputAccessionsPath", "./accessions.csv", "Path to output accessions csv file")
	outputLeanContextPath := flag.String("outputLCPath", "./LeanContext.csv", "Path to output LeanContext csv file")
	selection := ihec.PopulateFiles(*jsonPath)
	selection.PopulateAccessions()
	selection.PrintAccessions()
	selection.PrintLeanContext()
	selection.ExportAccessions(*outputAccessionsPath)
	selection.ExportLeanContext(*outputLeanContextPath)
}
