package main

import "fmt"

func checkRecord(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return false
	}

	var (
		countAbsent     int = 0
		countConsecLate int = 0
	)

	for _, value := range s {
		if value == 'A' {
			countAbsent++
			countConsecLate = 0
		} else if value == 'L' {
			countConsecLate++
			if countConsecLate == 3 {
				return false
			}
		} else if value == 'P' {
			countConsecLate = 0
		}
	}

	return countAbsent < 2 && countConsecLate != 3
}

func main() {
	fmt.Println(checkRecord("PPALLP"))
	fmt.Println(checkRecord("PPALLL"))
}
