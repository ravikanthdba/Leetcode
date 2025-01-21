package main

func Reverse(s string) string {
	var output string

	if len(s) <= 1 {
		output += s
	}

	return Reverse(s[1:]) + output
}

func main() {
	Reverse("abc")
}
