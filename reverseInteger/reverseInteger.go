package main

import (
	"fmt"
	"math"
	"strconv"
)

func power(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	if x > 0 && x < 10 {
		return x
	}

	count := 0

	absoluteValue := int(math.Abs(float64(x)))

	n := len(strconv.Itoa(absoluteValue)) - 1

	for x != 0 {
		rem := x % 10
		count += (rem * power(10, n))
		x = x / 10
		n--
	}

	fmt.Println(count)

	if x < 0 {
		return -1 * count
	}

	return count
}

func main() {
	fmt.Println(reverse(123))
}
