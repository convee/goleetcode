package main

// 236. 二叉树的最近公共祖先
// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
// 思路：后序遍历 DFS
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// 递归灵魂三问
// 1、这个函数是干嘛的？（该函数输入三个参数root，p，q，它会返回一个节点）
// 2、这个函数参数中的变量是什么？（函数参数中的变量是root，因为根据框架，lowestCommonAncestor(root)会递归调用root.left和root.right；至于p和q，我们要求它俩的公共祖先，它俩肯定不会变化的。）
// 3、得到函数的递归结果，你应该干什么？

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 如果 p、q 都在以 root 为根节点的树中，那么 left、right 分别是 p、q
	if left != nil && right != nil {
		return root
	}
	// 如果 p、q 都不在以 root 为根节点的树中，直接返回 nil
	if left == nil && right == nil {
		return nil

	}
	// 如果 p、q 只有一个存在 root 为根节点的树中，函数返回该节点
	if left == nil {
		return right
	}
	return left
}
