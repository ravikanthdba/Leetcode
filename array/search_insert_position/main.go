package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		if target > nums[0] {
			return len(nums) + 1
		}
		return len(nums) - 1
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] < target {
			continue
		} else {
			return i
		}
	}

	return len(nums)
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}
