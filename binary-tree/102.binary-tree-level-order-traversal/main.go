package main

import "container/list"

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

//dfs 递归 思路
//对该二叉树进行先序遍历（根左右的顺序），
//同时，记录节点所在的层次level，并且对每一层都定义一个数组，然后将访问到的节点值放入对应层的数组中。

func levelOrder2(root *TreeNode) [][]int {
	return dfs(root, 0, [][]int{})
}

func dfs(root *TreeNode, level int, res [][]int) [][]int {
	if root == nil {
		return res
	}
	if level == len(res) {
		res = append(res, []int{})
	} else {
		res[level] = append(res[level], root.Val)
	}
	res = dfs(root.Left, level+1, res)
	res = dfs(root.Right, level+1, res)
	return res
}

//BFS 思路：
//使用Queue的数据结构。我们将root节点初始化进队列，通过消耗尾部，插入头部的方式来完成BFS
func levelOrder3(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	// 定义一个双向队列
	queue := list.New()
	// 头部插入根节点
	queue.PushFront(root)
	// 进行广度搜索
	for queue.Len() > 0 {
		var current []int
		listLength := queue.Len()
		for i := 0; i < listLength; i++ {
			//消耗尾部
			// queue.Remove(queue.Back()).(*TreeNode)：移除最后一个元素并将其转化为TreeNode类型
			node := queue.Remove(queue.Back()).(*TreeNode)
			current = append(current, node.Val)
			if node.Left != nil {
				//插入头部
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
		}
		result = append(result, current)
	}
	return result
}
