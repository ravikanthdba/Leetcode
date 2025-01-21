package main

import "fmt"

func moveZeroes(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}

	if len(nums) == 2 {
		if nums[0] == 0 && nums[1] != 0 {
			nums[0], nums[1] = nums[1], nums[0]
			return nums
		} else if nums[0] != 0 && nums[1] == 0 {
			return nums
		} else if nums[0] != 0 && nums[1] != 0 {
			return nums
		} else if nums[0] == 0 && nums[1] == 0 {
			return nums
		}
	}

	p := 0
	q := len(nums) - 1
	for p < q {
		if nums[p] == 0 && nums[q] != 0 {
			nums[p], nums[q] = nums[q], nums[p]
			p++
			q--
		} else if nums[p] != 0 && nums[q] == 0 {
			p++
		} else if nums[p] != 0 && nums[q] != 0 {
			p++
		} else if nums[p] == 0 && nums[q] == 0 {
			q--
		}
	}

	fmt.Println(nums)

	var startingPoint int
	for key, value := range nums {
		if value == 0 {
			startingPoint = key
			break
		}
	}

	p = 0
	q = startingPoint - 1

	for p <= q {
		if nums[p] > nums[q] {
			nums[p], nums[q] = nums[q], nums[p]
			q--
		} else if nums[p] == nums[q] {
			p++
			q--
		} else if nums[p] < nums[q] {
			p++
		}
	}

	fmt.Println(nums)

	return nums
}

func main() {
	fmt.Println(moveZeroes([]int{0, 1, 0}))
	//fmt.Println(moveZeroes([]int{0}))
}
