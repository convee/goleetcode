package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 145. 二叉树的后序遍历
// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
// 思路：递归
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

// 思路：迭代
func postorderTraversal(root *TreeNode) (res []int) {
	stk := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}
	}
	return
}
