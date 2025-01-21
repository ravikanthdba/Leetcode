package main

func fun1(n int) int {
	if n > 100 {
		return n - 10
	} else {
		return fun1(fun1(n + 11))
	}
}

func main() {
	println(fun1(95))
}
