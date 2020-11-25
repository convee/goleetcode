package main

//92. 反转链表 II
//https://leetcode-cn.com/problems/reverse-linked-list-ii/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// 关键点是m的前一个节点
	dummy := &ListNode{
		Next: head,
	}
	prem := dummy
	for i := 1; i <= m-1; i++ {
		prem = prem.Next
	}

	// m和n区间内反转
	current := prem.Next
	var pre *ListNode
	for i := m; i <= n; i++ {
		tempNext := current.Next
		current.Next = pre
		pre = current
		current = tempNext
	}

	// prem.Next为m节点，原始m节点应该指向n+1节点
	prem.Next.Next = current
	// prem指向pre（反转之后的头节点）
	prem.Next = pre

	return dummy.Next
}
