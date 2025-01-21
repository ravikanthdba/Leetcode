package main

import "fmt"

func repeatedCharacter(s string) string {
	if len(s) == 0 {
		return ""
	}

	if len(s) == 1 {
		return ""
	}

	if len(s) == 2 {
		if s[0] == s[1] {
			return string(s[0])
		}
		return ""
	}

	var charString = make(map[string]int)

	for _, v := range s {
		charString[string(v)]++
	}

	for _, value := range s {
		if charString[string(value)] == 2 {
			return string(value)
		}
	}

	return ""
}

func main() {
	fmt.Println(repeatedCharacter("abccbaacz"))
}
