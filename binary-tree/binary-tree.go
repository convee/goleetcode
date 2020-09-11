package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	// 先访问根再访问左右
	fmt.Println(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{}
	preorderTraversal(root)
}
