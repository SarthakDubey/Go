package course1

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func MakeJsonHelper(name *string, address *string) {
	cache := map[string]string{"name": strings.TrimSpace(*name), "address": strings.TrimSpace(*address)}
	if data, err := json.Marshal(cache); err == nil {
		fmt.Println(string(data))
	}
}

func MakeJson() {
	fmt.Println("Enter the Name:")
	reader := bufio.NewReader(os.Stdin)
	name, err1 := reader.ReadString('\n')
	fmt.Println("Enter the Address:")
	address, err2 := reader.ReadString('\n')
	if err1 != nil || err2 != nil {
		log.Fatal(err1, err2)
	}
	MakeJsonHelper(&name, &address)
}
