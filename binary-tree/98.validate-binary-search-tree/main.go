package main

// 98. 验证二叉搜索树
// https://leetcode-cn.com/problems/validate-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validBST(root, nil, nil)
}

func validBST(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return validBST(root.Left, min, root) && validBST(root.Right, root, max)

}
