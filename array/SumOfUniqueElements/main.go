package main

import (
	"fmt"
	"sort"
)

func sumOfUnique(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		if nums[0] == nums[1] {
			return 0
		}
		return nums[0] + nums[1]
	}

	sort.Ints(nums)

	candidate := 0
	count := 0
	p := 0
	sum := 0

	for p < len(nums) {
		if candidate == nums[p] {
			count++
		} else {
			if count == 1 {
				sum += candidate
			}
			candidate = nums[p]
			count = 1
		}
		p++
	}

	if count > 1 {
		return sum
	}

	return sum + candidate
}

func main() {
	fmt.Println(sumOfUnique([]int{10, 4, 10, 9, 5}))
}
