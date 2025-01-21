package main

import "fmt"

func StaircaseTraversal(height int, maxSteps int) int {
	// Write your code here.
	if height == 0 || height == 1 {
		return 1
	}

	return -1
}

func main() {
	fmt.Println(StaircaseTraversal(4, 2))
}
