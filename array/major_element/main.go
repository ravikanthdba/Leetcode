package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		if nums[0] == nums[1] {
			return nums[0]
		}
		return 0
	}

	sort.Ints(nums)

	var (
		candidate        = -1
		count            = 1
		currentCandidate int
		currentCount     int
	)

	for _, value := range nums {
		if value == candidate {
			count++
		} else {
			if count >= currentCount {
				currentCandidate = candidate
				currentCount = count
			}
			candidate = value
			count = 1
		}
	}

	if count >= currentCount {
		currentCandidate = candidate
		currentCount = count
	}

	if currentCount > (len(nums) / 2) {
		return currentCandidate
	}

	return 0
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}
