package main

import "fmt"

var counter int = 0

func func1(x int) int {
	if x > 0 {
		return func1(x-1) + x
	}
	return 0
}

func func2(y int) int {
	if y > 0 {
		counter++
		return func2(y-1) + counter
	}
	return 0
}

func main() {
	x := 5
	println("counter = ", counter)
	fmt.Println(func1(x))
	fmt.Println(func2(x))
}
