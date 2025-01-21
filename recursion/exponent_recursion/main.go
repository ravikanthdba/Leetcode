package main

func function1(m, n int) int {
	if n == 0 {
		return 1
	}

	return m * function1(m, n-1)
}

func main() {
	println(function1(2, 3))
	println(function1(3, 3))
	println(function1(5, 3))
}
