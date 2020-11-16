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

// 思路：迭代
func inorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}
