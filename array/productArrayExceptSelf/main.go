package main

import "fmt"

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return nums
	}

	if len(nums) == 2 {
		nums[0], nums[1] = nums[1], nums[0]
		return nums
	}

	var (
		prefix_product []int
		prefix_pr      = 1
		suffix_product []int
		suffix_pr      = 1
		result         []int
	)

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			prefix_product = append(prefix_product, prefix_pr)
		} else {
			prefix_pr = prefix_product[len(prefix_product)-1]
			prefix_product = append(prefix_product, prefix_pr*nums[i-1])
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			suffix_product = append(suffix_product, suffix_pr)
		} else {
			suffix_pr = suffix_product[len(suffix_product)-1]
			suffix_product = append(suffix_product, suffix_pr*nums[i+1])
		}
	}

	for i := 0; i < len(prefix_product); i++ {
		result = append(result, prefix_product[i]*suffix_product[(len(suffix_product)-1)-i])
	}

	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
