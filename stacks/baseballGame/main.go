package main

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	if len(operations) == 0 {
		return 0
	}

	if len(operations) == 1 {
		if operations[0] == "C" {
			return 0
		} else if operations[0] == "D" {
			return 0
		} else {
			value, _ := strconv.Atoi(operations[0])
			return value
		}
	}

	var (
		stack []int
		top   = -1
	)

	for _, value := range operations {
		if value == "D" {
			if top == -1 {
				continue
			}

			result := stack[top] * 2
			stack = append(stack, result)
			top++
		} else if value == "+" {
			if top == -1 || top == 0 {
				continue
			}

			stack = append(stack, stack[top]+stack[top-1])
			top++
		} else if value == "C" {
			stack = stack[:top]
			top--
		} else {
			res, _ := strconv.Atoi(value)
			stack = append(stack, res)
			top++
		}

	}

	sum := 0
	for _, value := range stack {
		sum += value
	}

	return sum
}

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
}
