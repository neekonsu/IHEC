package ihec

// Browser represents the table located at the bottom of the IHEC page after data has been selected
type Browser struct {
	DataSelectorRows []string `css:"dataSelector__row"`
}
