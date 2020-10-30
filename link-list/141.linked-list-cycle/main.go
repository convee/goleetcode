package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//141. 环形链表
//https://leetcode-cn.com/problems/linked-list-cycle/
// 输入：head = [3,2,0,-4], pos = 1
// 输出：true
// 解释：链表中有一个环，其尾部连接到第二个节点。
// 思路：通过hash表来检测节点之前是否被访问过，来判断链表是否成环
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]int)
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = 1
		head = head.Next
	}
	return false
}

// 思路：双指针法
// 解题方法：通过使用具有 不同速度 的快、慢两个指针遍历链表，空间复杂度可以被降低至 O(1)。
// 慢指针每次移动一步，而快指针每次移动两步。

func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	for fast != nil && head != nil && fast.Next != nil {
		if fast == head {
			return true
		}
		fast = fast.Next.Next
		head = head.Next
	}
	return false
}
