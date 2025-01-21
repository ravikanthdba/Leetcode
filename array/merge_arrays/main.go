package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	if m == 0 && n == 0 {
		return []int{}
	}

	if m > 0 && n == 0 {
		return nums1
	}

	if m == 0 && n > 0 {
		for i := 0; i < len(nums2); i++ {
			nums1[i] = nums2[i]
		}
		return nums1
	}

	var (
		temp   = nums1[:m]
		result []int
		i      int
		j      int
	)

	for i < len(temp) && j < len(nums2) {
		if temp[i] < nums2[j] {
			result = append(result, temp[i])
			i++
		} else if temp[i] > nums2[j] {
			result = append(result, nums2[j])
			j++
		} else {
			result = append(result, temp[i])
			result = append(result, nums2[j])
			i++
			j++
		}
	}

	if i < len(temp) && j == len(nums2) {
		result = append(result, temp[i:]...)
		nums1 = result
	} else if i == len(temp) && j < len(nums2) {
		result = append(result, nums2[j:]...)
		nums1 = result
	}

	return nums1
}

func main() {
	fmt.Println(merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
}
