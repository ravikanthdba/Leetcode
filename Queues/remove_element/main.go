package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		if nums[0] == val {
			return 0
		}
		return 1
	}

	if len(nums) == 2 {
		if nums[0] == nums[1] && nums[0] == val {
			nums = []int{}
			fmt.Println(nums)
			return 0
		}

		if nums[0] != nums[1] && nums[0] == val {
			nums = nums[1:]
			fmt.Println(nums)
			return 1
		}

		if nums[0] != nums[1] && nums[1] == val {
			nums = nums[:1]
			fmt.Println(nums)
			return 1
		}

		return 2
	}

	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left] == val && nums[right] == val {
			right--
		} else if nums[left] == val && nums[right] != val {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		} else if nums[left] != val && nums[right] == val {
			left++
		} else {
			left++
			right--
		}
	}

	//if left >= right {
	//	return 0
	//}

	return len(nums[:left+1])
}

func main() {
	fmt.Println(removeElement([]int{4, 5}, 4))
}
