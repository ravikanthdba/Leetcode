package main

import "fmt"

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

	for p.Next != nil && q.Next != nil {
		if p.Val < q.Val {
			if mergedList == nil {
				mergedList = &ListNode{Val: p.Val}
			} else {
				current := mergedList.Next
				for current.Next != nil {
					current = current.Next
				}
				current.Next = &ListNode{Val: p.Val}
			}
			p = p.Next
		} else if p.Val > q.Val {
			if mergedList == nil {
				mergedList = &ListNode{Val: q.Val}
			} else {
				current := mergedList.Next
				for current.Next != nil {
					current = current.Next
				}
				current.Next = &ListNode{Val: q.Val}
			}
			q = q.Next
		} else if p.Val == q.Val {
			if mergedList == nil {
				mergedList = &ListNode{Val: p.Val}
				mergedList.Next = &ListNode{Val: q.Val}
			} else {
				current := mergedList.Next
				for current.Next != nil {
					current = current.Next
				}
				current.Next = &ListNode{Val: p.Val}
				current.Next.Next = &ListNode{Val: q.Val}
			}
			p = p.Next
			q = q.Next
		}
	}

	if p != nil && q == nil {
		mergedList.Next = p
	} else if p == nil && q != nil {
		mergedList.Next = q
	} else {
		mergedList.Next = p
		mergedList.Next = q
	}

	return mergedList
}

func main() {
	list1 := &ListNode{Val: 5, Next: nil}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	fmt.Println(mergeTwoLists(list1, list2))
}
