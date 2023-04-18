package course1

import "fmt"

func HelloHelper() string {
	return "Hello, World!"
}

func Hello() {
	fmt.Println(HelloHelper())
}
