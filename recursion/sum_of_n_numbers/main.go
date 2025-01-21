package main

func func1(n int) int {
	if n == 0 {
		return 0
	}

	return n + func1(n-1)
}

func main() {
	println(func1(5))
	// println(func1(10))
}
