package course3

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func userInput(nums *[]int) {
	fmt.Println("Enter integers separated with whitespace that need sorting. Hit Enter to begin")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	inputArray := strings.Split(input.Text(), " ")
	for _, val := range inputArray {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		*nums = append(*nums, num)
	}
}

func QuadrantSort(nums []int, wg *sync.WaitGroup, ch chan []int) {
	fmt.Println("Sorting integers: ", nums)
	sort.Ints(nums)
	ch <- nums
	wg.Done()
}

func MergeTwo(nums1, nums2 []int) []int {
	i, j, n, m, result := 0, 0, len(nums1), len(nums2), make([]int, 0)
	for i < n && j < m {
		if nums1[i] < nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}
	if i < n {
		result = append(result, nums1[i:]...)
	} else {
		result = append(result, nums2[j:]...)
	}
	return result
}

func SortIntegers() {
	nums := make([]int, 0)
	var wg sync.WaitGroup
	userInput(&nums)
	n := len(nums)
	quadrant, chunkSize, ch := 4, n/4, make(chan []int)
	for start := 0; start < n && quadrant > 0; start += chunkSize {
		end := min(n, start+chunkSize)
		if quadrant == 1 {
			end = n
		}
		quadrant--
		wg.Add(1)
		go QuadrantSort(nums[start:end], &wg, ch)
	}
	result := make([]int, 0)
	for quadrant < 4 {
		result = MergeTwo(result, <-ch)
		quadrant++
	}
	fmt.Printf("\nSorted result: %d\n", result)
	wg.Wait()
}
