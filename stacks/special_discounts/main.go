package main

import "fmt"

func finalPrices(prices []int) []int {
	if len(prices) == 0 {
		return []int{}
	}

	if len(prices) == 1 {
		return prices
	}

	if len(prices) == 2 {
		if prices[0] > prices[1] {
			return []int{prices[0] - prices[1], prices[1]}
		}
		return prices
	}

	var stack []int
	//var i int
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i] >= prices[j] {
				stack = append(stack, prices[i]-prices[j])
				break
			}
		}
	}

	if len(prices) != len(stack) {
		prices = prices[len(stack):]
		stack = append(stack, prices...)
	}

	return stack
}

func main() {
	fmt.Println(finalPrices([]int{4, 7, 1, 9, 4, 8, 8, 9, 4}))
}
