package main

// 203. 移除链表元素
// https://leetcode-cn.com/problems/remove-linked-list-elements/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	sentinel := &ListNode{0, head}
	pre := sentinel

	for head != nil {
		if head.Val == val {
			pre.Next = head.Next
		} else {
			pre = head
		}
		head = head.Next
	}
	return sentinel.Next

}
