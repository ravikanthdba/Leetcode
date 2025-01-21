package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		stack  []TreeNode
		result []int
	)

	for root != nil || len(stack) != 0 {
		if root != nil {
			result = append(result, root.Val)
			stack = append(stack, *root)
			root = root.Left
		} else {
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = p.Right
		}
	}

	return result
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		stack  []TreeNode
		result []int
	)

	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, *root)
			root = root.Left
		} else {
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, p.Val)
			root = p.Right
		}
	}

	return result
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		stack  []TreeNode
		result []int
	)

	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, *root)
			root = root.Left
		} else {
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, p.Val)
			if p.Right != nil {
				stack = append(stack, p)
				root = p.Right
				result = append(result, root.Val)
			}
		}
	}

	return result
}

func main() {
	t := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
	}

	fmt.Println(preorderTraversal(t))
	fmt.Println(inorderTraversal(t))
	fmt.Println(postorderTraversal(t))
}
