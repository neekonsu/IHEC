package main

import (
	"flag"
	"fmt"

	ihec "github.com/neekonsu/IHEC"
)

func main() {
	jsonPath := flag.String("jsonPath", "./res/json_metadata", "Path to directory storing JSON metadata files")
	// TODO: Make sure you name the downloaded HTML file as "page.html"
	// TODO: implement function to handle html: htmlPath := flag.String("htmlPath", "./res/page.html", "Path to html file of IHEC browser **after** selecting data of interest")
	selection := ihec.PopulateFiles(*jsonPath)
	selection.PopulateAccessions()
	for _, item := range selection.Accessions {
		fmt.Println(item)
	}
}
