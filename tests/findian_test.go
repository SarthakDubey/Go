package tests

import (
	"CourseraGo/course1"
	"strconv"
	"testing"
)

func TestCheck(t *testing.T) {
	var tests = []struct {
		inputString    string
		expectedOutput bool
	}{
		{"ian", true},
		{"Ian", true},
		{"iuiygaygn", true},
		{"I d skd a efju N", true},
		{"ihhhhhn", false},
		{"ina", false},
		{"xian", false},
	}
	for _, test := range tests {
		test := test
		t.Run(test.inputString, func(t *testing.T) {
			t.Parallel()
			result := course1.Check(test.inputString)
			if result != test.expectedOutput {
				t.Errorf("Expected %s, got %s", strconv.FormatBool(test.expectedOutput), strconv.FormatBool(result))
			}
		})
	}
}
