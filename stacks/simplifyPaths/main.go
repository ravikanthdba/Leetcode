package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {

	canonicalPathList := strings.Split(path, "/")

	if len(canonicalPathList) == 0 {
		return "/"
	}

	var (
		stack []string
		top   = -1
	)

	for _, value := range canonicalPathList {
		if value != "" {
			if value == ".." {
				if top == -1 {
					continue
				}

				stack = stack[:top]
				top--
			} else if value == "." {
				continue
			} else {
				stack = append(stack, value)
				top++
			}
		}
	}

	return "/" + strings.Join(stack, "/")
}

func main() {
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
}
