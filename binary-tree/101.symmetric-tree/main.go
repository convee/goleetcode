package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
// 101.对称二叉树
// https://leetcode-cn.com/problems/symmetric-tree/
// 思路：递归
// 1.两个树的跟节点相同
// 2.每个树的左子树和另外一个树的右子树相同
// 3.每个树的右子树和另外一个树的左子树相同
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	if !check(a.Left, b.Right) {
		return false
	}
	if !check(a.Right, b.Left) {
		return false
	}
	return true
}
