package tests

import (
	"CourseraGo/course1"
	"bufio"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestStreamReader(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     []course1.Person
	}{
		{name: "Testing sample data", filename: "data/first_last_name.txt", want: []course1.Person{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.Open(tt.filename)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				fields := strings.Split(line, " ")
				tt.want = append(tt.want, *course1.CreatePerson(fields[0], strings.Join(fields[1:], " ")))
			}
			if got := course1.StreamReader(bufio.NewScanner(file)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StreamReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
