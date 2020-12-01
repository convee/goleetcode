package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 116. 填充每个节点的下一个右侧节点指针
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/
// 思路：层次遍历（广度优先搜索）
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	//初始化队列，将根节点加入队列
	queue := []*Node{root}
	//循环迭代层数
	for len(queue) > 0 {
		curQueue := queue
		queue = nil
		//遍历这一层所有节点
		for i, node := range curQueue {
			//连接当前层所有节点
			if i < len(curQueue)-1 {
				node.Next = curQueue[i+1]
			}
			//拓展下一层节点
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

// 思路：递归
func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

func connectTwoNode(node1 *Node, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	connectTwoNode(node1.Left, node1.Right)
	connectTwoNode(node2.Left, node2.Right)
	connectTwoNode(node1.Right, node2.Left)
}
