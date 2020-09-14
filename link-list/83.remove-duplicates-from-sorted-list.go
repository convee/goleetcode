package main


//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次
func deleteDuplicates2(head *ListNode) *ListNode {
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

func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head                             // 终止条件
	}
	head.Next  = deleteDuplicates3(head.Next)     // 递归调用
	if head.Val == head.Next.Val {
		head = head.Next                      // 去重
	}
	return head
}
