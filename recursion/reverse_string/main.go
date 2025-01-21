package main

import "fmt"

var output string

func Reverse(n string) string {
	if len(n) == 0 {
		return ""
	}

	if len(n) == 1 {
		return output + n
	}

	Reverse(n[1:])

	return output
}

func main() {
	fmt.Println(Reverse("Hello"))
}
