package main

import "fmt"

var n = []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

func Fibonacci(number int) int {
	if number <= 1 {
		n[number] = number
		return number
	}

	if n[number-2] == -1 {
		n[number-2] = Fibonacci(number - 2)
	}

	if n[number-1] == -1 {
		n[number-1] = Fibonacci(number - 1)
	}

	return Fibonacci(number-2) + Fibonacci(number-1)
}

func main() {
	fmt.Println(Fibonacci(5))
}
