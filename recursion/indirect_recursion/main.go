package main

func fn1(n int) {
	if n > 0 {
		println(n)
		fn2(n - 1)
	}
}

func fn2(n int) {
	if n > 1 {
		println(n)
		fn1(n / 2)
	}
}

func main() {
	fn1(20)
}
