package main

import (
	"flag"

	ihec "github.com/neekonsu/IHEC"
)

func main() {
	jsonPath := flag.String("jsonPath", "./res", "Path to directory storing JSON metadata files")
	agreementAccessionsPath := flag.String("agreementPath", "./agreementAccessions.csv", "Path to agreement accessions csv")
	outputAccessionsPath := flag.String("outputAccessionsPath", "./accessions.csv", "Path to output accessions csv file")
	outputLeanContextPath := flag.String("outputLCPath", "./LeanContext.csv", "Path to output LeanContext csv file")
	outputIntersectionsPath := flag.String("outputIntersectionsPath", "./intersections.csv", "Path to output intersecting accessions")
	selection := ihec.PopulateFiles(*jsonPath)
	selection.PopulateAccessions()
	selection.PrintLeanContext()
	selection.PrintAccessions()
	selection.ExportAccessions(*outputAccessionsPath)
	selection.ExportLeanContext(*outputLeanContextPath)
	ihec.ExportIntersections(*outputAccessionsPath, *agreementAccessionsPath, *outputIntersectionsPath)
}
