package main

// 230. 二叉搜索树中第K小的元素
// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
// 思路：
// 1、对于 BST 的每一个节点node，左子树节点的值都比node的值要小，右子树节点的值都比node的值大。
// 2、对于 BST 的每一个节点node，它的左侧子树和右侧子树都是 BST。
// 3、BST 的中序遍历结果是有序的（升序）。
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
