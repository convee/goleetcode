package main

import (
	"goleetcode/util"
)
type ListNode struct {
	Val  int
	Next *ListNode
}

//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		// 全部删除完再移动到下一个元素
		for current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return head
}

func main() {
	var head *util.ListNode
	util.FmtListNode(head)
}