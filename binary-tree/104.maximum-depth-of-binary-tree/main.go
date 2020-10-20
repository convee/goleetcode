package main

//104. 二叉树的最大深度
//https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

// 二叉树的深度计算，首先要判断节点，以下是计算二叉树的详细步骤：
// 1、一颗树只有一个节点，它的深度是1；
// 2、二叉树的根节点只有左子树而没有右子树，那么可以判断，二叉树的深度应该是其左子树的深度加1；
// 3、二叉树的根节点只有右子树而没有左子树，那么可以判断，那么二叉树的深度应该是其右树的深度加1；
// 4、二叉树的根节点既有右子树又有左子树，那么可以判断，那么二叉树的深度应该是其左右子树的深度较大值加1。
// 一棵深度为k，且有2^k-1个节点的二叉树，称为满二叉树。这种树的特点是每一层上的节点数都是最大节点数。
// 而在一棵二叉树中，除最后一层外，若其余层都是满的，并且最后一层或者是满的，或者是在右边缺少连续若干节点，则此二叉树为完全二叉树。
// 具有n个节点的完全二叉树的深度为floor(log2n)+1。深度为k的完全二叉树，至少有2k-1个叶子节点，至多有2k-1个节点。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
