package main

import (
	"fmt"
	"sort"
)

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 && nums[0] == val {
		return 0
	}

	sort.Ints(nums)

	start_point_location := 0
	end_point_location := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == val {
			start_point_location = i
			break
		} else {
			continue
		}
	}

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == val {
			end_point_location = i
			break
		} else {
			continue
		}
	}

	fmt.Println(start_point_location)
	fmt.Println(end_point_location)

	return 0
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
