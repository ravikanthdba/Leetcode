package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {
	if len(tickets) == 0 {
		return 0
	}

	if k > len(tickets)-1 {
		return 0
	}

	front := 0
	time := 0

	for tickets[k] != 0 {
		if tickets[front] != 0 {
			tickets[front]--
			if front < len(tickets)-1 {
				front++
			} else {
				front = 0
			}
			time++
		} else {
			if front < len(tickets)-1 {
				front++
			} else {
				front = 0
			}
			continue
		}
	}

	return time
}

func main() {
	fmt.Println(timeRequiredToBuy([]int{5, 1, 1, 1}, 0))
}
