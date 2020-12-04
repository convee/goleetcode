package main

import "strconv"

// 652. 寻找重复的子树
// https://leetcode-cn.com/problems/find-duplicate-subtrees/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	myMap := make(map[string]int)
	ret := make([]*TreeNode, 0)

	dnf(root, myMap, &ret)
	return ret
}

func dnf(root *TreeNode, myMap map[string]int, ret *[]*TreeNode) string {
	if root == nil {
		return ""
	}

	lstring := dnf(root.Left, myMap, ret)
	rstring := dnf(root.Right, myMap, ret)

	retstring := lstring + "," + rstring + "," + strconv.Itoa(root.Val)

	if v, ok := myMap[retstring]; !ok {
		myMap[retstring] = 1
	} else {
		if v == 1 {
			*ret = append(*ret, root)
		}
		myMap[retstring]++
	}

	return retstring
}
