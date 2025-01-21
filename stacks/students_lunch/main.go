package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	if len(students) == 0 && len(sandwiches) != 0 {
		return 0
	}

	if len(students) != 0 && len(sandwiches) == 0 {
		return len(students)
	}

	count := 0

	for len(sandwiches) != 1 && count != len(students) {
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			count = 0
		} else {
			top := len(students) - 1
			temp := students[top]
			for top != 0 {
				students[top] = students[top-1]
				top--
			}
			students[0] = temp
			count++
		}
	}

	if count == len(students) {
		return len(students)
	}

	if students[0] == sandwiches[0] {
		return 0
	}

	return len(students)
}

func main() {
	fmt.Println(countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}))
}
