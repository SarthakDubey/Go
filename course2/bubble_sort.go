package course2

import (
	"fmt"
)

func inputToSlice(unsorted []int, err error, counts *int) []int {
	if err != nil {
		return unsorted
	}
	if *counts == 10 {
		fmt.Println("Got maximum allowed 10 numbers, discarding others.")
		return unsorted
	}
	var num int
	count, err := fmt.Scan(&num)
	if count == 1 {
		unsorted = append(unsorted, num)
	}
	*counts += count
	return inputToSlice(unsorted, err, counts)
}

func Swap(slice []int, index int) {
	slice[index], slice[index+1] = slice[index+1], slice[index]
}

func BubbleSort(unsorted []int) {
	didSwap := true
	for didSwap {
		didSwap = false
		for i := 0; i < len(unsorted)-1; i++ {
			if unsorted[i] > unsorted[i+1] {
				didSwap = true
				Swap(unsorted, i)
			}
		}
	}
}

func BubbleSortUserInput() {
	fmt.Println("Enter a sequence of integers (maximum 10 allowed):")
	counts := 0
	unsorted := inputToSlice([]int{}, nil, &counts)
	fmt.Printf("Unsorted list of numbers entered: %v\n", unsorted)
	BubbleSort(unsorted)
	fmt.Printf("Sorted list of numbers entered: %v\n", unsorted)
}
