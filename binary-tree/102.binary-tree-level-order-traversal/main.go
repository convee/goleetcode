package main

//102. 二叉树的层序遍历
//https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//思路：
//广度优先搜索
//BFS
//1.首先根元素入队
//2.当队列不为空的时候
//	1>求当前队列的长度 Si
//	2>依次从队列中取 Si 个元素进行拓展，然后进入下一次迭代
//它和 BFS 的区别在于 BFS 每次只取一个元素拓展，而这里每次取 Si 个元素。在上述过程中的第 i 次迭代就得到了二叉树的第 i 层的 Si 个元素
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}
