package main

import "fmt"

func isValid(s string) bool {
	var bracketsMap = map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
		"]": "[",
		"}": "{",
		")": "(",
	}

	if len(s) == 0 || len(s) == 1 || len(s)%2 == 0 {
		return false
	}

	if len(s) == 2 {
		return string(bracketsMap[string(s[0])]) == string(s[1])
	}

	leftPtr := 0
	rightPtr := len(s) - 1

	for leftPtr < rightPtr {
		if string(bracketsMap[string(s[leftPtr])]) == string(s[rightPtr]) {
			leftPtr++
			rightPtr--
		} else if string(bracketsMap[string(s[leftPtr])]) == string(s[leftPtr+1]) && leftPtr <= rightPtr {
			leftPtr += 2
		} else if string(bracketsMap[string(s[rightPtr])]) == string(s[rightPtr-1]) && leftPtr <= rightPtr {
			rightPtr -= 2
		} else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isValid("(){}}{"))
}
