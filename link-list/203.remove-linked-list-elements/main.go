package main

// 203. 移除链表元素
// https://leetcode-cn.com/problems/remove-linked-list-elements/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路：双指针
func removeElements(head *ListNode, val int) *ListNode {
	// 定义哨兵结点
	sentinel := &ListNode{0, head}
	pre := sentinel //前指针
	cur := head     //当前指针
	for cur != nil {
		// 前指针移动
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		// 当前指针移动
		cur = cur.Next
	}
	return sentinel.Next
}
