package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n > 0 {
		return x * myPow(x, n-1)
	}

	return x * (1 / myPow(x, -1*n))
}

func main() {
	fmt.Println(myPow(2, -8))
}
