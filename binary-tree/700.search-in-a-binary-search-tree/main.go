package main

// 700. 二叉搜索树中的搜索
// https://leetcode-cn.com/problems/search-in-a-binary-search-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	}
	if val > root.Val {
		return searchBST(root.Right, val)
	}
	return root

}
