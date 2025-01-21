package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 0 {
		return false
	}

	if n == 0 {
		return true
	}

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) &&
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
	}

	return n == 0
}

func main() {
	flowerbed := []int{1, 0, 1, 0, 1, 0, 1}
	n := 1
	fmt.Println(canPlaceFlowers(flowerbed, n))
}
