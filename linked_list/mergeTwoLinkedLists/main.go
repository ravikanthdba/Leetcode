package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil && list2 != nil {
		return list2
	}

	if list2 == nil && list1 != nil {
		return list1
	}

	p := list1
	q := list2

	var (
		mergedList *ListNode
	)

	for p != nil && q != nil {
		if p.Val > q.Val {
			if mergedList == nil {
				mergedList = &ListNode{Val: q.Val, Next: nil}
			} else {
				current := mergedList
				for current.Next != nil {
					current = current.Next
				}
				current.Next = &ListNode{Val: q.Val, Next: nil}
			}
			q = q.Next
		} else if p.Val < q.Val {
			if mergedList == nil {
				mergedList = &ListNode{Val: p.Val, Next: nil}
			} else {
				current := mergedList
				for current.Next != nil {
					current = current.Next
				}
				current.Next = &ListNode{Val: p.Val, Next: nil}
			}
			p = p.Next
		} else if p.Val == q.Val {
			if mergedList == nil {
				mergedList = &ListNode{Val: p.Val, Next: nil}
				mergedList.Next = &ListNode{Val: q.Val, Next: nil}
			} else {
				current := mergedList
				for current.Next != nil {
					current = current.Next
				}
				temp := &ListNode{Val: p.Val, Next: nil}
				temp.Next = &ListNode{Val: q.Val, Next: nil}
				current.Next = temp
			}
			p = p.Next
			q = q.Next
		}
	}

	if p != nil && q == nil {
		current := mergedList
		for current.Next != nil {
			current = current.Next
		}
		current.Next = p
	} else if p == nil && q != nil {
		current := mergedList
		for current.Next != nil {
			current = current.Next
		}
		current.Next = q
	} else {
		current := mergedList
		for current.Next != nil {
			current = current.Next
		}
		current.Next = p
		current.Next = q
	}

	return mergedList
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	//list1 := &ListNode{Val: 1, Next: nil}
	//list2 := &ListNode{Val: 2, Next: nil}

	merged := mergeTwoLists(list1, list2)

	current := merged
	for current != nil {
		fmt.Println("current: ", current.Val)
		current = current.Next
	}
}
