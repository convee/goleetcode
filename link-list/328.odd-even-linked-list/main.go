package main

// 328. 奇偶链表
// https://leetcode-cn.com/problems/odd-even-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路：
// 结点1作为奇链表的头，结点2作为偶链表的头
// 从第3个结点开始，依次轮流附在奇偶链的后面
// 遍历完后，奇链表的尾指向偶链表的头
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	o := head      //o 为奇链表尾结点
	p := head.Next //p 为偶链表头结点
	e := p         //e 为偶链表尾结点
	for o.Next != nil && e.Next != nil {
		o.Next = e.Next
		o = o.Next
		e.Next = o.Next
		e = e.Next

	}
	o.Next = p
	return head
}
