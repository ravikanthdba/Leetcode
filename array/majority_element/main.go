package main

import "fmt"

func isMajorityElement(nums []int, target int) bool {
	var elementCount = make(map[int]int)

	for _, val := range nums {
		elementCount[val]++
	}

	max := -99999
	var maxElement int
	for key, val := range elementCount {
		if val > max {
			max = val
			maxElement = key
		}
	}

	if max > len(nums)/2 && maxElement == target {
		return true
	}

	return false
}

func main() {
	fmt.Println(isMajorityElement([]int{10, 100, 101, 101}, 101))
}
