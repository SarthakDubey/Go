package tests

import (
	"CourseraGo/assesments"
	"testing"
)

func TestTruncHelper(t *testing.T) {
	type args struct {
		input float64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Test 1233.234", args{1233.234}, 1233},
		{"Test -12.34", args{-12.34}, -12},
		{"Test -0.34", args{-0.34}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := assesments.TruncHelper(tt.args.input); got != tt.want {
				t.Errorf("TruncHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}
