package main

import (
	"fmt"
)

  type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 

func searchBST(root *TreeNode, val int) *TreeNode {
	if root != nil {
		if root.Val == val {
			fmt.Println("root is equal", root)
			return root
		}

		if val < root.Val {
			return searchBST(root.Left, val)
		}

		if val > root.Val {
			return searchBST(root.Right, val)
		}
	}
    
    return nil
}

func main() {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}

	value := 2

	fmt.Println(searchBST(tree, value))
}