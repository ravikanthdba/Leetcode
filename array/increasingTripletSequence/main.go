package main

import "fmt"

func increasingTriplet(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 || len(nums) == 2 {
		return false
	}

	r := 0
	q := r + 1
	p := q + 1

	for p <= len(nums)-1 {
		if nums[r] < nums[q] && nums[q] < nums[p] {
			return true
		} else {
			r = q
			q = p
			p++
		}
	}

	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}
