package main

// 652. 寻找重复的子树
// https://leetcode-cn.com/problems/find-duplicate-subtrees/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []*TreeNode
var m map[string]int

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m = map[string]int{}
	subtree(root)
	return res
}

func subtree(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	left := subtree(root.Left)
	right := subtree(root.Right)
	str := left + "," + right + "," + string(root.Val)
	if _, ok := m[str]; ok {
		m[str]++
	} else {
		m[str] = 1
	}
	if m[str] == 2 {
		res = append(res, root)
	}

	return str

}
