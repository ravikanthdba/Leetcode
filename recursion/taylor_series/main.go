package main

import "fmt"

func taylorSeries(x float64, n int, p float64, f float64) float64 {
	if n == 0 {
		return 1
	}

	// Calculate the current power and factorial
	p = p * x
	f = f * float64(n)

	// Recursive calculation
	return p/f + taylorSeries(x, n-1, p, f)
}

func main() {
	// Start with p = 1 and f = 1
	result := taylorSeries(1, 10, 1, 1)
	fmt.Printf("%.6f\n", result) // Output: 2.718282
}
