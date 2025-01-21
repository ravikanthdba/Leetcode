package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var (
		array  []int
		output *ListNode
	)

	for head.Next != nil {
		array = append(array, head.Val)
		head = head.Next
	}

	array = append(array, head.Val)

	for i := len(array) - 1; i >= 0; i-- {
		if i == len(array)-1 {
			output = &ListNode{Val: array[i], Next: nil}
		} else {
			if output.Next == nil {
				output.Next = &ListNode{Val: array[i], Next: nil}
			} else {
				current := output.Next
				for current.Next != nil {
					current = current.Next
				}
				current.Next = &ListNode{Val: array[i], Next: nil}
			}
		}
	}

	return output
}

func main() {
	ll := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	reverseList(ll)
}
