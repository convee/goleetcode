package util

import (
	"container/list"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
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