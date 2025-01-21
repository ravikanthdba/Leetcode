package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 && len(nums2) == 0 {
		return []int{}
	}

	if (len(nums1) != 0 && len(nums2) == 0) || (len(nums1) == 0 && len(nums2) != 0) {
		return []int{}
	}

	var (
		hashNums1 = make(map[int]int)
		hashNums2 = make(map[int]int)
		result    []int
	)

	for _, value := range nums1 {
		hashNums1[value]++
	}

	for _, value := range nums2 {
		hashNums2[value]++
	}

	for key, value := range hashNums1 {
		if val, ok := hashNums2[key]; ok {
			if val > value {
				for i := 0; i < value; i++ {
					result = append(result, key)
				}
			} else if val < value {
				for i := 0; i < val; i++ {
					result = append(result, key)
				}
			} else if val == value {
				for i := 0; i < val; i++ {
					result = append(result, key)
				}
			}
		} else {
			continue
		}
	}

	return result
}

func main() {
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
