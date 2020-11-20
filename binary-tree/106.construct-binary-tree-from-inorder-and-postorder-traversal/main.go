package main

// 106. 从中序与后序遍历序列构造二叉树
// 根据一棵树的中序遍历与后序遍历构造二叉树。
// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

// 中序遍历 inorder = [9,3,15,20,7]
// 后序遍历 postorder = [9,15,7,20,3]
// 3
// / \
// 9  20
//  /  \
// 15   7

// 思路：递归
// 后序遍历 最后一个元素为根节点
// 中序遍历 已知根节点 分成左子树、右子树

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
