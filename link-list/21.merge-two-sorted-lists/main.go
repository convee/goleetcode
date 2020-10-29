package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//21. 合并两个有序链表
//https://leetcode-cn.com/problems/merge-two-sorted-lists/
//思路：
//首先我们维护一个 prehead 的哨兵节点。我们其实只需要调整它的 next 指针。让它总
// 是指向l1或者 l2 中较小的一个，直到 l1 或者 l2 任一指向null。这样到了最后，如果 l1 还是 l2 中任意
// 一方还有余下元素没有用到，那余下的这些元素一定大于 prehead 已经合并完的链表（因为是有序链表）。
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{} //哨兵节点
	result := prehead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prehead.Next = l1
			l1 = l1.Next
		} else {
			prehead.Next = l2
			l2 = l2.Next
		}
		prehead = prehead.Next
	}
	if l1 != nil {
		prehead.Next = l1
	}
	if l2 != nil {
		prehead.Next = l2
	}
	return result.Next
}
