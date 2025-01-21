package main

import "fmt"

func isValidParenthesis(s string) bool {
	return s == "(" || s == ")"
}

func maxDepth(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	var (
		parenthesis     = map[string]string{")": "("}
		stack           []string
		top             = -1
		depthMax, depth int
	)

	for _, value := range s {
		if isValidParenthesis(string(value)) {
			_, ok := parenthesis[string(value)]
			if ok {
				stack = stack[:top]
				top--
				if depth > depthMax {
					depthMax = depth
				}
			} else {
				stack = append(stack, string(value))
				top++
			}
			depth = len(stack)
		}
	}

	return depthMax
}

func main() {
	fmt.Println(maxDepth("1+(2*3)/(2-1)"))
}
