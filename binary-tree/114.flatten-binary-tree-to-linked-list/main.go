package main

// 114. 二叉树展开为链表
// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var last *TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	flatten(root.Left)

	root.Right = last
	root.Left = nil
	last = root

}
