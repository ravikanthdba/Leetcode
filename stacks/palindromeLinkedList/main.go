package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var (
		stack []int
		top   = -1
	)

	if head.Next == nil {
		return true
	}

	for head != nil {
		if top == -1 {
			stack = append(stack, head.Val)
			top++
		} else {
			if stack[top] == head.Val {
				stack = stack[:top]
				top--
			} else {
				stack = append(stack, head.Val)
				top++
			}
		}
		head = head.Next
	}

	if len(stack) == 0 {
		return true
	}

	return false
}

func main() {
	linkedList := ListNode{Val: 1}
	linkedList.Next = &ListNode{Val: 0}
	linkedList.Next.Next = &ListNode{Val: 1}
	fmt.Println(isPalindrome(&linkedList))
}
