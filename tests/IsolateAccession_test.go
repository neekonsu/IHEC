package tests

import (
	"testing"

	ihec "github.com/neekonsu/IHEC"
)

func TestIsolateAccession(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Testing IsolateAccession with example URL",
			args: args{"https://www.ebi.ac.uk/ega/datasets/EGAD00001003963"},
			want: "EGAD00001003963",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ihec.IsolateAccession(tt.args.path); got != tt.want {
				t.Errorf("IsolateAccession() = %v, want %v", got, tt.want)
			}
		})
	}
}
