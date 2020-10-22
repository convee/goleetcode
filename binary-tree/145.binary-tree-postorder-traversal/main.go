package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//145. 二叉树的后序遍历
//https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
func postorderTraversal(root *TreeNode) []int {
	var res []int
	return postorder(root, res)

}
func postorder(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = postorder(root.Left, res)
	res = postorder(root.Right, res)
	res = append(res, root.Val)
	return res
}
