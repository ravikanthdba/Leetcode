package main

import (
	"fmt"
	"strings"
)

func removeDuplicates(s string) string {
	if len(s) == 0 {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	if len(s) == 2 {
		if s[0] == s[1] {
			return ""
		}
		return s
	}

	var (
		stack []string
		top   = -1
	)

	for _, value := range s {
		if top == -1 {
			stack = append(stack, string(value))
			top++
		} else if stack[top] == string(value) {
			stack = stack[:top]
			top--
		} else {
			stack = append(stack, string(value))
			top++
		}
	}

	return strings.Join(stack, "")
}

func main() {
	fmt.Println(removeDuplicates("azxxzy"))
}
