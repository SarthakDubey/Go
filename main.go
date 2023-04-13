package main

import (
	"CourseraGo/assesments"
)

//func isValid(s string) bool {
//	cache := map[rune]rune{'(': ')', '{': '}', '[': ']'}
//	var stack []rune
//	for _, rune := range s {
//		if rune == '(' || rune == '{' || rune == '[' {
//			stack = append(stack, cache[rune])
//		} else {
//			if len(stack) == 0 {
//				return false
//			}
//			index := len(stack) - 1
//			popped := stack[index]
//			stack = (stack)[:index]
//			if popped != rune {
//				return false
//			}
//		}
//	}
//	return true
//}

func main() {
	//assesments.Hello()
	//assesments.Trunc()
	//assesments.Findian()
	//assesments.Slice()
	assesments.MakeJson()
}
