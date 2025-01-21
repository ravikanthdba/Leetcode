package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 && len(t) == 0 {
		return false
	}

	if len(s) > len(t) {
		return false
	}

	if len(s) == 1 && len(t) == 1 {
		if s == t {
			return true
		}
		return false
	}

	p := 0
	q := 0

	for p < len(s)-1 && q < len(t)-1 {
		if s[p] == t[q] {
			p++
			q++
		} else {
			q++
		}
	}

	if p == len(s)-1 && q != len(t)-1 {
		if s[p] == t[q] {
			return true
		}

		for q < len(t) {
			if s[p] == t[q] {
				return true
			}
			q++
		}

		return false
	} else if p == len(s)-1 && q == len(t)-1 {
		if s[p] == t[q] {
			return true
		}
		return false
	}

	return false
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
