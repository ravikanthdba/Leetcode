package main

import (
	"fmt"
	"math"
)

func convertDigitsToNumber(digits []int) int {
	sum := 0
	j := int(math.Pow(10, float64(len(digits)-1)))
	for i := 0; i < len(digits); i++ {
		sum += digits[i] * j
		j = j / 10
	}
	return sum
}

func splitToDigits(number int) []int {
	var digits []int

	if number < 10 {
		digits = []int{number}
		return digits
	}

	for number != 0 {
		remainder := number % 10
		digits = append(digits, remainder)
		number = number / 10
	}

	var reversedDigits []int
	for i := len(digits) - 1; i >= 0; i-- {
		reversedDigits = append(reversedDigits, digits[i])
	}
	return reversedDigits
}

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return nil
	}

	number := convertDigitsToNumber(digits) + 1
	return splitToDigits(number)
}

func main() {
	//fmt.Println(plusOne([]int{1, 2, 3}))
	//fmt.Println(plusOne([]int{9, 9}))
	//fmt.Println(plusOne([]int{0}))
	fmt.Println(plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}))
}
