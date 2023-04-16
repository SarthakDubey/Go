package assesments

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func CreatePerson(fname string, lname string) *Person {
	person := new(Person)
	person.fname, person.lname = fname, lname
	return person
}

func StreamReader(filebuffer *bufio.Scanner) []Person {
	persons := make([]Person, 0)
	for filebuffer.Scan() {
		line := filebuffer.Text()
		fields := strings.Split(line, " ")
		persons = append(persons, Person{fields[0], strings.Join(fields[1:], " ")})
	}
	return persons
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// ReadFileFirstLastNameFile Reads a file address as user input
// ReadFileFirstLastNameFile Validate file exists!
func ReadFileFirstLastNameFile() {
	fmt.Println("Enter filename to parse first, and last names:")
	var filename string
	_, err := fmt.Scanln(&filename)
	check(err)
	if fileExists(filename) {
		file, err := os.Open(filename)
		check(err)
		defer file.Close()
		for _, person := range StreamReader(bufio.NewScanner(file)) {
			fmt.Printf("First name: %s, Last name: %s\n", person.fname, person.lname)
		}
	}
}
