package main

// 160. 相交链表
// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路：哈希表，时间复杂度 O(m+n)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	m := make(map[*ListNode]int)
	for headA != nil {
		m[headA] = 1
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := m[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil

}

// 思路：暴力解法
