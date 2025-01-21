package main

import (
	"fmt"
	"sort"
)

type TwoSum struct {
	Nums []int
}

func Constructor() TwoSum {
	return TwoSum{
		Nums: []int{},
	}
}

func (this *TwoSum) Add(number int) {
	this.Nums = append(this.Nums, number)
}

func (this *TwoSum) Find(value int) bool {
	if len(this.Nums) == 0 {
		return false
	}

	if len(this.Nums) == 1 {
		return false
	}

	if len(this.Nums) == 2 {
		return this.Nums[0]+this.Nums[1] == value
	}

	sort.Ints(this.Nums)
	p := 0
	q := len(this.Nums) - 1

	for p < q {
		if this.Nums[p]+this.Nums[q] < value {
			p++
		} else if this.Nums[p]+this.Nums[q] > value {
			q--
		} else {
			return true
		}
	}

	return false
}

func main() {
	t := Constructor()
	t.Add(0)
	t.Add(0)
	fmt.Println(t.Find(0))
}
