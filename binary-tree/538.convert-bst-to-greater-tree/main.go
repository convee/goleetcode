package main

// 538. 把二叉搜索树转换为累加树
// https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	var sum int
	doConvertBST(root, &sum)
	return root
}

func doConvertBST(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	doConvertBST(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	doConvertBST(root.Left, sum)
}
