package main

// 234. 回文链表
// https://leetcode-cn.com/problems/palindrome-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	vars := []int{}
	for head != nil {
		vars = append(vars, head.Val)
		head = head.Next
	}
	n := len(vars)
	for i := 0; i < (n+1)/2; i++ {
		if vars[i] != vars[n-1-i] {
			return false
		}
	}
	return true
}
