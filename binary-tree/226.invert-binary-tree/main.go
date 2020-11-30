package main

// 226. 翻转二叉树
// https://leetcode-cn.com/problems/invert-binary-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：先序遍历
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	right := root.Right
	root.Right = root.Left
	root.Left = right
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
