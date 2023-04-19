package tests

import (
	"CourseraGo/course2"
	"math"
	"testing"
)

func TestGenDisplaceFn(t *testing.T) {
	type args struct {
		a  float64
		v0 float64
		s0 float64
		t  float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Testing 10, 2, 1", args{10.0, 2.0, 1.0, 3.0}, 52.0},
		{"Testing 10, 2, 1", args{10.0, 2.0, 1.0, 5.0}, 136.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := course2.GenDisplaceFn(&tt.args.a, &tt.args.v0, &tt.args.s0)
			got := fn(&tt.args.t)
			if math.Abs(tt.want-got) > 0.0001 {
				t.Errorf("GenDisplaceFn() = %.2f, want %.2f", got, tt.want)
			}
		})
	}
}
