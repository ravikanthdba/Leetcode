package main

import "fmt"

func contains(word string, letter string) int {
	var position int
	for key, value := range word {
		if letter == string(value) {
			position = key
			return position
		}
	}
	return -1
}

func reversePrefix(word string, ch byte) string {
	if len(word) == 0 {
		return ""
	}

	if len(word) == 1 || len(word) == 2 {
		return word
	}

	position := contains(word, string(ch))

	if position == -1 {
		return word
	}

	var (
		stack  []string
		output string
		top    = -1
	)

	part1 := word[:position+1]
	part2 := word[position+1:]

	for part1 != "" {
		stack = append(stack, string(part1[0]))
		top++
		part1 = part1[1:]
	}

	for top != -1 {
		output += stack[top]
		top--
	}

	return output + part2
}

func main() {
	fmt.Println(reversePrefix("abcdefd", 'd'))
}
