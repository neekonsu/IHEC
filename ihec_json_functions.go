package ihec

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// CheckErr is a generic error handler that can be used for unimportant errors
func CheckErr(msg string, err error) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

// ParseJSON takes the filepath of the JSON metadata and returns a fully unmarshalled METADATA object
func ParseJSON(path string) METADATA {
	var output METADATA
	reader, err := os.Open(path)
	CheckErr("Unable to open file at "+path+": ", err)
	decoder := json.NewDecoder(reader)
	decoder.Decode(&output)
	return output
}

/*
IsolateAccession takes a RawDataURL (string) and returns an isolated EGA (accession) number (string)
Takes url such as: "https://www.ebi.ac.uk/ega/datasets/EGAD00001003963"
*/
func IsolateAccession(path string) string {
	URL, err := url.Parse(path)
	CheckErr("Unable to parse URL: "+path+": ", err)
	index := strings.Index(URL.Path, "datasets/") + 9
	return URL.Path[index:]
}

// PopulateFiles takes a path to a directory storing JSON metadata files and returns a Selection with files but not Accessions
func PopulateFiles(path string) Selection {
	var output Selection
	CheckErr("Unable to list directory: ",
		filepath.Walk(path, func(path string, fileInfo os.FileInfo, err error) error {
			output.files = append(output.files, ParseJSON(path))
			if err != nil {
				return err
			}
			return nil
		}))
	return output
}
