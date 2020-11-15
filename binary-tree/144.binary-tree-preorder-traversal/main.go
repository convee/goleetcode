package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 144. 二叉树的前序遍历
// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

// 思路：递归
func preorderTraversal(root *TreeNode) []int {
	var res []int
	res = preorder(root, res)
	return res
}
func preorder(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = preorder(root.Left, res)
	res = preorder(root.Right, res)
	return res
}

// 思路：迭代
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for 0 < len(stack) || root != nil { //root != nil 只为了第一次root判断，必须放最后
		for root != nil {
			res = append(res, root.Val)       //前序输出
			stack = append(stack, root.Right) //右节点 入栈
			root = root.Left                  //移至最左
		}
		index := len(stack) - 1 //栈顶
		root = stack[index]     //出栈
		stack = stack[:index]
	}
	return res
}
