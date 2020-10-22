package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//94. 二叉树的中序遍历
//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
func inorderTraversal(root *TreeNode) []int {
	var res []int
	res = inorder(root, res)
	return res
}

func inorder(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = inorder(root.Left, res)
	res = append(res, root.Val)
	res = inorder(root.Right, res)
	return res
}
