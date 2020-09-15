package main
import (
	"container/list"
	"fmt"
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

func FmtListNode(head *ListNode) {
	l := list.New()
	for ; head != nil; head = head.Next {
	   l.PushFront(head.Val)
	}
	for item := l.Front(); item != nil; item = item.Next() {
	   fmt.Println(item.Value)
	}
 }
