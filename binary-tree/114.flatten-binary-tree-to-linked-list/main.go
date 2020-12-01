package main

// 114. 二叉树展开为链表
// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 定义：将以 root 为根的树拉平为链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	flatten(root.Left)
	/**** 后序遍历位置 ****/
	// 1、左右子树已经被拉平成一条链表
	left := root.Left
	right := root.Right
	// 2、将左子树作为右子树
	root.Right = left
	root.Left = nil
	// 3、将原先的右子树接到当前右子树的末端
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right

}
