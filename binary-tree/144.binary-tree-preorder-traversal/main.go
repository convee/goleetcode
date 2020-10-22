package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//144. 二叉树的前序遍历
//https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
func preorderTraversal(root *TreeNode) []int {
	var res []int
	res = dfs(root, res)
	return res
}
func dfs(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = dfs(root.Left, res)
	res = dfs(root.Right, res)
	return res
}
