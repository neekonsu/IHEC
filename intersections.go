package ihec

import (
	"encoding/csv"
	"os"
)

// ExportIntersections takes paths to two n*1 matricies of accessions and returns the intersections between them.
func ExportIntersections(path1 string, path2 string, path3 string) {
	file1, err := os.Open(path1)
	CheckErr("Couldn't open file: ", err)
	file2, err := os.Open(path2)
	CheckErr("Couldn't open file: ", err)

	reader1 := csv.NewReader(file1)
	reader2 := csv.NewReader(file2)

	csv1, err := reader1.ReadAll()
	CheckErr("Couldn't read file: ", err)
	csv2, err := reader2.ReadAll()
	CheckErr("Couldn't read file: ", err)

	set1 := matrixToSlice(csv1)
	set2 := matrixToSlice(csv2)

	commonAccessions := intersection(set1, set2)
	table := sliceToMatrix(commonAccessions)

	output, err := os.Create(path3)
	CheckErr("Unable to initialize new file: ", err)
	writer := csv.NewWriter(output)
	defer writer.Flush()
	CheckErr("Unable to write csv: ", writer.WriteAll(table))
}

// transpose column and return as slice
func matrixToSlice(matrix [][]string) []string {
	output := make([]string, len(matrix))
	for i, item := range matrix {
		output[i] = item[0]
	}
	return output
}

// transpose slice and return as column
func sliceToMatrix(slice []string) [][]string {
	output := make([][]string, len(slice))
	for i, item := range slice {
		output[i] = []string{item}
	}
	return output
}

// copied from stackOverflow in the interest of time
func intersection(a []string, b []string) []string {
	var output []string
	for _, i := range a {
		for _, j := range b {
			if i == j {
				output = append(output, j)
			}
		}
	}
	return output
}
