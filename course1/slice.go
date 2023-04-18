package course1

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func InsertAndSortSlice(num int, nums *[]int) {
	if *nums != nil {
		*nums = append(*nums, num)
		sort.Ints(*nums)
	} else {
		log.Panic("Empty slice provided.")
	}
}

func Slice() {
	nums := make([]int, 0)
	var userInput string
	for {
		fmt.Print("Enter an integer to sort, X to end: ")
		_, err := fmt.Scanln(&userInput)

		if err != nil {
			fmt.Println("Invalid Input: ", err)
		}

		if strings.ToLower(userInput) == "x" {
			fmt.Println("Exiting...")
			os.Exit(0)
		}
		if num, err := strconv.Atoi(userInput); err == nil {
			InsertAndSortSlice(num, &nums)
		}
		fmt.Println(nums)
	}
}
