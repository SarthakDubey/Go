package assesments

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Check(str string) bool {
	str = strings.ToLower(str)
	return strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a")
}

func Findian() {
	fmt.Print("Enter the string to scan: ")
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	if Check(strings.TrimSpace(userInput)) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
