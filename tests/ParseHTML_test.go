package tests

import (
	"reflect"
	"testing"

	ihec "github.com/neekonsu/IHEC"
)

func TestParseHTML(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want ihec.Browser
	}{
		{
			name: "Generic test for this function",
			args: args{path: "./res/page.html"},
			want: ihec.Browser{DataSelectorRows: []string{"", ""}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ihec.ParseHTML(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
