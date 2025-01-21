package main

import (
	"fmt"
	"strings"
)

func contains(wordsList []string, word string) bool {
	for _, value := range wordsList {
		if value == word {
			return true
		}
	}

	return false
}

func minLength(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	subStrings := []string{"AB", "CD"}

	if len(s) == 2 {
		if contains(subStrings, s) {
			return 0
		} else {
			return 2
		}
	}

	var stack []string
	var top int = -1

	for i := 0; i < len(s); i++ {
		if top < 0 {
			// if stack is empty
			stack = append(stack, string(s[i]))
			top++
		} else if top == 0 {
			// if stack has only 1 element
			stack = append(stack, string(s[i]))
			top++
		} else {
			if !contains(subStrings, stack[top-1]+stack[top]) {
				// if stack[top - 1] + stack[top] - concatenation of last two is not in substring
				stack = append(stack, string(s[i]))
				top++
			} else {
				stack = stack[:len(stack)-2]
				top -= 2
				stack = append(stack, string(s[i]))
				top++
			}
		}
	}

	//if len(strings.Join(stack, "")) == 2 {
	//	if contains(subStrings, strings.Join(stack, "")) {
	//		return 0
	//	}
	//}

	if len(stack) > 1 {
		if contains(subStrings, stack[top-1]+stack[top]) {
			return len(strings.Join(stack, "")) - 2
		}
	}

	return len(strings.Join(stack, ""))
}

func main() {
	fmt.Println(minLength("ABG"))
}
