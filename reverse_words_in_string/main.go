package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	fmt.Println(words, len(words))
	var output string

	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != "" {
			if i > 0 {
				output += strings.TrimSpace(words[i]) + " "
			} else {
				output += strings.TrimSpace(words[i])
			}
		}
	}

	return strings.TrimSpace(output)
}

func main() {
	s := "a good   example"
	println(reverseWords(s))
}
