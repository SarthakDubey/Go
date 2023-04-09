package tests

import (
	"CourseraGo/assesments"
	"reflect"
	"testing"
)

func TestInsertAndSortSlice(t *testing.T) {
	type args struct {
		num  int
		nums *[]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Testing 1", args{1, &[]int{}}, []int{1}},
		{"Testing 2", args{2, &[]int{1}}, []int{1, 2}},
		{"Testing 0", args{0, &[]int{1, 2}}, []int{0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			nums := make([]int, len(*tt.args.nums))
			copy(nums, *tt.args.nums)
			assesments.InsertAndSortSlice(tt.args.num, &nums)
			if !reflect.DeepEqual(nums, tt.want) {
				t.Errorf("InsertAndSortSlice() = %v, want %v", *tt.args.nums, tt.want)
			}
		})
	}
}
