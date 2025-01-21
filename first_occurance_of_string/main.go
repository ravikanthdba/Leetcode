package main

func strStr(haystack string, needle string) int {
	if haystack == "" {
		return -1
	}

	if needle == "" {
		return -1
	}

	if len(haystack) < len(needle) {
		return -1
	}

	return index[0]
}

func main() {
	haystack := "sadbutsad"
	needle := "sad"
	println(strStr(haystack, needle))
}
