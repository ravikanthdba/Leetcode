package main

import "fmt"

type NumArray struct {
	Nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		Nums: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if len(this.Nums) == 0 {
		return 0
	}

	if len(this.Nums) == 1 {
		if left == right && left == 0 {
			return this.Nums[0]
		}
		return 0
	}

	if left < 0 {
		left = 0
	}

	if right > len(this.Nums)-1 {
		right = len(this.Nums) - 1
	}

	sum := 0

	for left <= right {
		sum += this.Nums[left]
		left++
	}

	return sum
}

func main() {
	t := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(t.SumRange(0, 2))
}
