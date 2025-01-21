package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}

	if len(s) == 1 {
		return true
	}

	originalStringWithoutSpaces := ""
	for _, value := range s {
		if !unicode.IsPunct(value) && string(value) != " " && string(value) != "`" {
			originalStringWithoutSpaces += strings.ToLower(string(value))
		}
	}

	palindromeStringWithoutSpaces := ""
	for i := len(s) - 1; i >= 0; i-- {
		if !unicode.IsPunct(rune(s[i])) && string(s[i]) != " " && string(s[i]) != "`" {
			palindromeStringWithoutSpaces += strings.ToLower(string(s[i]))
		}
	}

	fmt.Println(originalStringWithoutSpaces)
	fmt.Println(palindromeStringWithoutSpaces)
	return originalStringWithoutSpaces == palindromeStringWithoutSpaces
}

func main() {
	fmt.Println(isPalindrome("`l;`` 1o1 ??;l`"))
}
