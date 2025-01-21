package main

import "fmt"

func contains(list []int, key int) bool {
	for _, value := range list {
		if value == key {
			return true
		}
	}
	return false
}

func isPossibleToSplit(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums)%2 != 0 {
		return false
	}

	var (
		res1 []int
		res2 []int
	)

	for _, value := range nums {
		if len(res1) == 0 && len(res2) == 0 {
			res1 = append(res1, value)
		} else if !contains(res1, value) && len(res1) <= len(res2) {
			res1 = append(res1, value)
		} else if !contains(res2, value) && len(res1) >= len(res2) {
			res2 = append(res2, value)
		} else {
			if !contains(res2, value) {
				temp := res2[len(res2)-1]
				res2[len(res2)-1] = res1[len(res1)-1]
				res1[len(res1)-1] = temp
				res2 = append(res2, value)
			} else if !contains(res1, value) {
				temp := res1[len(res1)-1]
				res1[len(res1)-1] = res2[len(res2)-1]
				res2[len(res2)-1] = temp
				res1 = append(res1, value)
			} else {
				continue
			}
		}
	}
	fmt.Println(res1)
	fmt.Println(res2)

	if len(nums)/2 == len(res1) && len(nums)/2 == len(res2) {
		return true
	}

	return false
}

func main() {
	fmt.Println(isPossibleToSplit([]int{8, 9, 8, 5, 9, 3, 3, 1, 2, 1}))
	fmt.Println(isPossibleToSplit([]int{1, 1, 2, 2, 3, 4}))
	fmt.Println(isPossibleToSplit([]int{10, 1, 7, 4, 5, 1, 6, 4}))
}
