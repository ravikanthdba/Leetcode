package main

func fn1(n int) {
	if n > 0 {
		println(n)
		fn1(n - 1)
		fn1(n - 1)
	}
}

func main() {
	fn1(3)
}
