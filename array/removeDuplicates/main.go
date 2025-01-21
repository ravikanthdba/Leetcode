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

	p_pos := 0
	p := 0
	q := p + 1

	for p < len(nums) {
		if nums[q] == nums[p] {
			p = q
			q++
		} else {
			p_pos = p
			for i := q; i < len(nums); i++ {
				nums[p] = nums[q]
				p++
			}
			p = p_pos
		}
	}

	return len(nums)
}

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
