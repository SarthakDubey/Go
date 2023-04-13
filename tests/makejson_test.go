package tests

import (
	"CourseraGo/assesments"
	"testing"
)

func Test_makeJson(t *testing.T) {
	type args struct {
		name    string
		address string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test case 1", args{"John Doe", "123 Main St, Anytown USA"}},
		{"Test case 2", args{"Jane Doe", ""}},
		{"Test case 3", args{"", ""}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assesments.MakeJsonHelper(&tt.args.name, &tt.args.address)
		})
	}
}
