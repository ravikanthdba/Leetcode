package main

import "fmt"

func containsVowel(s rune) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	for _, value := range vowels {
		if value == s {
			return true
		}
	}

	return false
}

func reverseVowels(s string) string {
	runes := []rune(s)

	if len(runes) == 0 || len(runes) == 1 {
		return s
	}

	if len(runes) == 2 {
		if containsVowel(runes[0]) && containsVowel(runes[1]) {
			runes[0], runes[1] = runes[1], runes[0]
			return string(runes)
		} else {
			return s
		}
	}

	lptr := 0
	rptr := len(runes) - 1
	// var updatedString string

	for lptr < rptr {
		if containsVowel(runes[lptr]) && containsVowel(runes[rptr]) {
			runes[lptr], runes[rptr] = runes[rptr], runes[lptr]
			lptr++
			rptr--
			fmt.Println("runes: ", string(runes))
		} else if !containsVowel(runes[lptr]) && containsVowel(runes[rptr]) {
			lptr++
		} else if containsVowel(runes[lptr]) && !containsVowel(runes[rptr]) {
			rptr--
		} else {
			lptr++
			rptr--
		}
	}

	return string(runes)
}

func main() {
	s := "IceCreAm"
	println(reverseVowels(s))
}
