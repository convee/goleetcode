package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//19. 删除链表的倒数第N个节点
//https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//搞个头结点
	node := &ListNode{0, head}
	// 双指针
	prev, pn := node, node
	// 第二个指针先完后移n，那么第二个指针移到底的时候，第一个指针就是我们要删除的节点
	// 两个指针同时向后移动
	for i := 1; i <= n+1; i++ {
		pn = pn.Next
	}
	// 第二个指针没到底之前，两个指针一起后移
	for pn != nil {
		prev = prev.Next
		pn = pn.Next
	}
	// 删除第一个指针对应的next
	prev.Next = prev.Next.Next
	// 输出啦
	return node.Next
}
