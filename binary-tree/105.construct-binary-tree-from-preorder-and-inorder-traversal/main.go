package main

// 105. 从前序与中序遍历序列构造二叉树
// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 思路：递归
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 || len(inorder) <= 0 {
		return nil
	}
	rv := preorder[len(preorder)-1]
	var i int
	for i = range inorder {
		if inorder[i] == rv {
			break
		}
	}
	return &TreeNode{
		Val:   rv,
		Left:  buildTree(preorder[:i], inorder[:i]),
		Right: buildTree(preorder[i:len(preorder)-1], inorder[i+1:]),
	}
}
