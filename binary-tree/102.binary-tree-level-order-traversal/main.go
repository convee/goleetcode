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
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		queue = p
	}
	return ret
}

//递归
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	res = dfs(root, 0, res)
	return res
}

func dfs(root *TreeNode, level int, res [][]int) [][]int {
	if root == nil {
		return res
	}
	if level == len(res) {
		res = append(res, []int{})
	}
	res[level] = append(res[level], root.Val)
	res = dfs(root.Left, level+1, res)
	res = dfs(root.Right, level+1, res)
	return res
}
