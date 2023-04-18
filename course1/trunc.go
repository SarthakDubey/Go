package course1

import "fmt"
import "log"

func TruncHelper(input float64) int64 {
	return int64(input)
}

func Trunc() {
	var sample float64
	fmt.Println("Enter the floating point number")
	_, err := fmt.Scan(&sample)
	if err != nil {
		log.Print("Invalid Input:", err)
	}
	fmt.Println("Integer of input:", TruncHelper(sample))
}
