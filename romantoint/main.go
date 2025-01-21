package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	if s == "" {
		return 0
	}

	var roman = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	literals := strings.Split(s, "")

	sum := 0

	if len(literals) == 0 {
		return 0
	}

	if len(literals) == 1 {
		return roman[literals[0]]
	}

	if len(literals) == 2 {
		if roman[literals[0]] < roman[literals[1]] {
			return roman[literals[1]] - roman[literals[0]]
		} else {
			return roman[literals[0]] + roman[literals[1]]
		}
	}

	p := 1
	q := 0

	for p < len(literals) {
		if roman[literals[q]] < roman[literals[p]] {
			sum += roman[literals[p]] - roman[literals[q]]
			q += 2
			p += 2
		} else {
			sum += roman[literals[q]]
			p++
			q++
		}

	}

	if q == len(literals)-1 {
		sum += roman[literals[q]]
	}

	return sum
}

func main() {
	fmt.Println(romanToInt("III"))

	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
