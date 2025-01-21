package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	if len(candies) == 0 {
		return nil
	}

	max := -1
	for _, v := range candies {
		if v > max {
			max = v
		}
	}

	var result []bool
	if extraCandies == 0 {
		for _, kid := range candies {
			if kid == max {
				result = append(result, true)
			} else {
				result = append(result, false)
			}
		}
	} else {
		for _, kid := range candies {
			if kid+extraCandies >= max {
				result = append(result, true)
			} else {
				result = append(result, false)
			}
		}
	}
	return result
}

func main() {
	candies := []int{4, 2, 1, 1, 2}
	extraCandies := 1
	fmt.Println(kidsWithCandies(candies, extraCandies))
}
