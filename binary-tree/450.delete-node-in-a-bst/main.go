package main

// 450. 删除二叉搜索树中的节点
// https://leetcode-cn.com/problems/delete-node-in-a-bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key == root.Val {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		min := getMin(root.Right)
		root.Val = min.Val
		root.Right = deleteNode(root.Right, min.Val)
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

// BST 最左边的就是最小的

func getMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}
