package main

import "fmt"

func GetPermutations(array []int) [][]int {
	if len(array) == 0 {
		return [][]int{}
	}

	var finalArray [][]int

	if len(array) == 1 {
		var output [][]int
		value := array[0]
		output = append(output, []int{value})
		return output
	}

	if len(array) == 2 {
		finalArray = append(finalArray, []int{array[0], array[1]})
		// finalArray = append(finalArray, []int{array[1], array[0]})
		return finalArray
	}

	for i := 0; i < len(array); i++ {
		var subArray []int
		subArray = append(subArray, array[:i]...)
		subArray = append(subArray, array[i+1:]...)
		permutations := GetPermutations(subArray)
		// for j := 0; j < len(permutations); j++ {
		// 	finalArray = append(finalArray, append([]int{array[i]}, permutations[j]...))
		// }
		finalArray = append(finalArray, permutations...)
	}

	return finalArray
}

func main() {
	fmt.Println(GetPermutations([]int{1, 2, 3, 4}))
}
