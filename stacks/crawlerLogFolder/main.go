package main

import "fmt"

func minOperations(logs []string) int {
	if len(logs) == 0 {
		return 0
	}

	if len(logs) == 1 {
		if logs[0] == "../" {
			return 0
		}
		return 1
	}

	var (
		stack []string
		top   = -1
	)

	for _, value := range logs {
		if value == "../" {
			if len(stack) == 0 {
				continue
			}
			stack = stack[:top]
			top--
		} else if value == "./" {
			continue
		} else {
			stack = append(stack, value)
			top++
		}
	}

	return len(stack)
}

func main() {
	fmt.Println(minOperations([]string{"./", "wz4/", "../", "mj2/", "../", "../", "ik0/", "il7/"}))
}
