package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

//前序递归
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	// 先访问根再访问左右
	fmt.Println(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

// V3：前序非递归遍历
func preorderTraversal2(root *TreeNode) []int {
	// 非递归
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			// 前序遍历，所以先保存结果
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

func postorderTraversal(root *TreeNode) []int {
	// 通过lastVisit标识右子节点是否已经弹出
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 这里先看看，先不弹出
		node := stack[len(stack)-1]
		// 根节点必须在右节点弹出之后，再弹出
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1] // pop
			result = append(result, node.Val)
			// 标记当前这个节点已经弹出过
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	preorderTraversal(root)
}
