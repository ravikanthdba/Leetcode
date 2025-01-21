package main

import (
	"fmt"
	"strings"
)

func backspaceCompare(s string, t string) bool {
	if len(s) == 0 && len(t) == 0 {
		return true
	}

	if len(s) == 1 && len(t) == 1 {
		if string(s[0]) == string(t[0]) {
			return true
		}
		return false
	}

	if (len(s) == 0 && len(t) != 0) || (len(s) != 0 && len(t) == 0) {
		return false
	}

	var (
		stackS []string
		stackT []string
		topS   = -1
		topT   = -1
	)

	for _, value := range s {
		if string(value) == "#" {
			if len(stackS) == 0 {
				continue
			}

			stackS = stackS[:topS]
			topS--
		} else {
			stackS = append(stackS, string(value))
			topS++
		}
	}

	for _, value := range t {
		if string(value) == "#" {
			if len(stackT) == 0 {
				continue
			}

			stackT = stackT[:topT]
			topT--
		} else {
			stackT = append(stackT, string(value))
			topT++
		}
	}

	return strings.Join(stackS, "") == strings.Join(stackT, "")
}

func main() {
	fmt.Println(backspaceCompare("ab##", "c#d#"))
}
