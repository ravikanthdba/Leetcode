package main

import (
	"fmt"
	"reflect"
)

func contains(array []int, set [][]int) bool {
	for _, value := range set {
		if reflect.DeepEqual(array, value) {
			return true
		}
	}
	return false
}

func Powerset(array []int) [][]int {
	// Write your code here.
	if len(array) == 0 {
		return [][]int{}
	}

	var finalArray [][]int

	if len(array) == 2 {
		if !contains([]int{}, finalArray) {
			finalArray = append(finalArray, []int{})
		}

		if !contains([]int{array[0]}, finalArray) {
			finalArray = append(finalArray, []int{array[0]})
		}

		if !contains([]int{array[1]}, finalArray) {
			finalArray = append(finalArray, []int{array[1]})
		}

		if !contains(array, finalArray) {
			finalArray = append(finalArray, array)
		}
	}

	for i := 0; i < len(array); i++ {
		var sub []int
		sub = append(sub, array[:i]...)
		sub = append(sub, array[i+1:]...)
		subsets := Powerset(sub)
		for j := 0; j < len(subsets); j++ {
			if !contains(subsets[j], finalArray) {
				finalArray = append(finalArray, subsets[j])
			}
		}
	}

	if !contains(array, finalArray) {
		finalArray = append(finalArray, array)
	}

	return finalArray
}

func main() {
	fmt.Println(Powerset([]int{1, 2, 3, 4}))
}
