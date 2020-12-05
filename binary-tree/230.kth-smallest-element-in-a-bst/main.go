package main

// 230. 二叉搜索树中第K小的元素
// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var rank int
var res int

func kthSmallest(root *TreeNode, k int) int {
	doKthSmallest(root, k)
	return res
}

func doKthSmallest(root *TreeNode, k int) {
	if root == nil {
		return
	}
	doKthSmallest(root.Left, k)
	rank++
	if rank == k {
		res = root.Val
		return
	}
	doKthSmallest(root.Left, k)
}
