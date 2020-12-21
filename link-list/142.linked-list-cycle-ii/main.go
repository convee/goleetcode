package main

// 142. 环形链表 II
// https://leetcode-cn.com/problems/linked-list-cycle-ii/
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func detectCycle2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

func detectCycle3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head
	for fast != slow {
		fast, slow = fast.Next, slow.Next
	}

	return fast
}
