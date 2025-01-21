package main

import (
	"fmt"
	"math"
)

func isPowerOfTwo(n int) bool {
	return int(math.Sqrt(float64(n)))%2 == 0
}

func main() {
	fmt.Println()
}
