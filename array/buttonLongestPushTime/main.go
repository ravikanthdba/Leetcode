package main

import (
	"fmt"
	"math"
)

func buttonWithLongestTime(events [][]int) int {
	if len(events) == 0 {
		return 0
	}

	if len(events) == 1 {
		return events[0][0]
	}

	if len(events) == 2 {
		if (events[1][1] - events[0][1]) > events[0][1] {
			return events[1][0]
		} else if (events[1][1] - events[0][1]) == events[0][1] {
			if events[1][0] > events[0][0] {
				return events[0][0]
			}
		}
		return events[1][0]
	}

	var (
		buttonTime [][]int
		i          int
	)

	for i = 0; i < len(events); i++ {
		if i == 0 {
			buttonTime = append(buttonTime, []int{events[i][0], events[i][1]})
		} else {
			buttonTime = append(buttonTime, []int{events[i][0], events[i][1] - events[i-1][1]})
		}
	}

	minTime := math.MinInt
	minKey := -1
	for _, val := range buttonTime {
		if val[1] >= minTime {
			if val[1] == minTime {
				if val[0] < minKey {
					minKey = val[0]
				} else {
					continue
				}
			} else {
				minTime = val[1]
			}
		} else {
			continue
		}
	}

	return minKey
}

func main() {
	fmt.Println(buttonWithLongestTime([][]int{
		[]int{18, 10},
		[]int{11, 20},
		//[]int{3, 9},
		//[]int{1, 15},
	}))
}
