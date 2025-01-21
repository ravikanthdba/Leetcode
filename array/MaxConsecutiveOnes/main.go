package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		if nums[0] == 1 {
			return 1
		}
		return 0
	}

	if len(nums) == 2 {
		if nums[0] == nums[1] && nums[0] == 1 {
			return 2
		} else if nums[0] != nums[1] {
			return 1
		}
		return 0
	}

	current_count := 0
	count := 0
	p := 0

	for p < len(nums) {
		if nums[p] == 1 {
			count++
		} else {
			if count > current_count {
				current_count = count
			}
			count = 0
		}
		p++
	}

	if count > current_count {
		current_count = count
	}

	return current_count
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}
