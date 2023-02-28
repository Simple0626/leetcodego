package main

import (
	"fmt"
)

func isValid(s string) bool {

	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, i := range s {
		if (i == '[') || (i == '(') || (i == '{') {
			stack = append(stack, i)
		} else if ((i == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			((i == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((i == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}

	}
	return len(stack) == 0

}

func main() {
	s := "([)]"
	fmt.Println(isValid(s))
}
