package main

import (
	"fmt"
	"strings"
)

func removeOuterParentheses(s string) string {
	if len(s) == 0 {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	var (
		//validParenthesis = map[string]string{")": "("}
		//stack            []string
		top    = 0
		output []string
		//res              = ""
	)

	// Iterate over each character in the string
	for _, ch := range s {
		// If it's the first opening parenthesis at depth 0, skip it
		if ch == '(' {
			if top > 0 {
				output = append(output, string(ch))
			}
			top++
		} else if ch == ')' {
			top--
			// If it's the closing parenthesis at a depth greater than 0, add it to the output
			if top > 0 {
				output = append(output, string(ch))
			}
		}
	}

	return strings.Join(output, "")
}

func main() {
	fmt.Println(removeOuterParentheses("(()(()))"))
}
