package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 11.路径之和
// https://leetcode-cn.com/problems/path-sum/
// 思路：递归

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	if hasPathSum(root.Left, sum-root.Val) {
		return true
	}
	if hasPathSum(root.Right, sum-root.Val) {
		return true
	}
	return false
}
