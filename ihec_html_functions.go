package ihec

import (
	"os"

	html "github.com/zlepper/encoding-html"
)

// ParseHTML takes the path to an HTML file and returns a Browser for that data
func ParseHTML(path string) Browser {
	var browser Browser
	file, err := os.Open(path)
	CheckErr("Couldn't open file "+path+": ", err)
	err = html.NewDecoder(file).Decode(&browser)
	CheckErr("Couldn't decode HTML: ", err)
	return browser
}
