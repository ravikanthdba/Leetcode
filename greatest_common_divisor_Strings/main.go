package main

import (
	"fmt"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	return str1[:gcd(len(str1), len(str2))]
}

func main() {
	str1 := "ABABAB"
	str2 := "AB"
	fmt.Println(gcdOfStrings(str1, str2))
}
