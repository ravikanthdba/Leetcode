package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return []string{string(nums[0])}
	}

	p := 1
	var (
		output []string
		q      string
	)

	for p < len(nums) {
		if q == "" && nums[p]-nums[p-1] == 1 {
			q = strconv.Itoa(nums[p-1])
			p++
		} else if nums[p]-nums[p-1] == 1 && q != "" {
			p++
		} else {
			if q != "" {
				q = q + "->" + strconv.Itoa(nums[p-1])
			} else {
				q = q + strconv.Itoa(nums[p-1])
			}

			output = append(output, q)
			q = ""
			p++
		}
	}

	if nums[p-1]-nums[p-2] == 1 && q != "" {
		q = q + "->" + strconv.Itoa(nums[p-1])
		output = append(output, q)
	} else {
		output = append(output, strconv.Itoa(nums[p-1]))
	}

	return output
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
