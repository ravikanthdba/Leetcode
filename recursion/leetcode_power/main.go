package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 1 / (x * myPow(x, (n*-1)-1))
	}

	// If x is very small, consider it as zero to avoid underflow
	if x < 1e-10 {
		return 0
	}

	return x * myPow(x, n-1)
}

func main() {
	fmt.Println(myPow(0.00001, 2147483647))
}
