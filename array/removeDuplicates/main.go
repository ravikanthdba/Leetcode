package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 0 {
		return 1
	}

	if len(nums) == 2 {
		if nums[0] == nums[1] {
			return 1
		}
		return 2
	}

	q := 0
	p := 1
	count := 0

	for p < len(nums) {
		if nums[q] == nums[p] {
			p++
			count++
		} else {
			if q == 0 {
				q = p - 1
				nums = nums[q:]
				
			} else {
				var temp []int
				temp = nums[:q+1]
				q = p - 1
				temp = append(temp, nums[q:])
				nums = temp
			}
		}
	}
	

	return len(nums)
}

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
