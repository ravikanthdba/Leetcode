package main

func fibonacci(n int) {
	if n == 0 {
		println(0)
	} else if n == 1 {
		println(1)
	} else {
		a := 0
		b := 1
		for i := 2; i <= n; i++ {
			c := a + b
			a = b
			b = c
		}
		println(b)
	}
}

func main() {
	fibonacci(5)
	fibonacci(10)
}
