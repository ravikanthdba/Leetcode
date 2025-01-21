package main

import "fmt"

type Position struct {
	count    int
	position []int
}

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}

	if len(s) == 1 {
		return 0
	}

	if len(s) == 2 {
		if s[0] == s[1] {
			return -1
		}
		return 0
	}

	front := 0
	rear := front + 1

	for front != len(s)-1 {
		if s[front] == s[rear] && rear != len(s)-1 {
			for s[front] == s[rear] {
				front++
				rear = front + 1
			}
			front++
			rear = front + 1
		} else {
			if rear < len(s)-1 {
				rear++
			} else if rear == len(s)-1 {
				if s[rear] == s[front] {
					return -1
				} else {
					return front
				}
			} else {
				return front
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))

}
