package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func listToNumber(l *ListNode) int {
	i := 0
	sum := 0
	for l != nil {
		sum = sum + (l.Val * int(math.Pow10(i)))
		i++
		l = l.Next
	}
	return sum
}

func (l *ListNode) Add(element int) *ListNode {
	if l == nil {
		l = &ListNode{Val: element, Next: nil}
		return l
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	var output = &ListNode{Val: 0, Next: nil}

	if l1 == nil && l2 != nil {
		number := listToNumber(l2)
		for number != 0 {
			rem := number % 10
			output.Next = &ListNode{Val: rem, Next: nil}
			number = number / 10
		}
		return output
	}

	if l1 != nil && l2 == nil {
		number := listToNumber(l1)
		for number != 0 {
			rem := number % 10
			output.Next = &ListNode{Val: rem, Next: nil}
			number = number / 10
			output = output.Next
		}
		return output
	}

	val1 := listToNumber(l1)
	val2 := listToNumber(l2)
	sum := val1 + val2

	for sum != 0 {
		rem := sum % 10
		output.Next = &ListNode{Val: rem, Next: nil}
		sum = sum / 10
	}
	return output
}

func main() {
	l1 := ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	fmt.Println(addTwoNumbers(&l1, nil))
}
