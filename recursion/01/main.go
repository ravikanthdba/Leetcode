package main

import "fmt"

func func1(x int) {
	if x > 0 {
		fmt.Println(x)
		func1(x - 1)
	}
}

func func2(x int) {
	if x > 0 {
		func2(x - 1)
		fmt.Println(x)
	}
}

func main() {
	func1(3)
	fmt.Println("\n\n")
	func2(3)
}
