package tests

import (
	"CourseraGo/assesments"
	"testing"
)

func TestHelloHelper(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Testing Hello World!", "Hello, World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := assesments.HelloHelper(); got != tt.want {
				t.Errorf("HelloHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}
