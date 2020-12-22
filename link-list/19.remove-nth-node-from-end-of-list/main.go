package main

type ListNode struct {
	Val  int
	Next *ListNode
}

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

//首先我们将添加一个哑结点作为辅助，该结点位于列表头部。哑结点用来简化某些极端情况，
//例如列表中只含有一个结点，或需要删除列表的头部。在第一次遍历中，我们找出列表的长度 LL。
//然后设置一个指向哑结点的指针，并移动它遍历列表，直至它到达第 (L - n)(L−n) 个结点那里。
//我们把第 (L - n)(L−n) 个结点的 next 指针重新链接至第 (L - n + 2)(L−n+2) 个结点，完成这个算法。

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	node := &ListNode{0, head}
	l := 0
	first := head
	for first != nil {
		l++
		first = first.Next
	}
	l -= n
	first = node
	for l > 0 {
		l--
		first = first.Next
	}
	first.Next = first.Next.Next
	return node.Next
}

//思路：
//通过前节点删除当前节点 pre.Next=pre.Next.Next
//先定义一个哨兵节点 result （比如删除头节点时，需要知道前一个节点）
//遍历链表 head节点，当遍历到第n个时，开始移动 pre 节点
//head 为nil时，pre便是倒数第n个节点的前节点
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	result := &ListNode{}
	result.Next = head
	var pre *ListNode
	cur := result
	i := 1
	for head != nil {
		if i >= n {
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}
	pre.Next = pre.Next.Next
	return result.Next
}

func removeNthFromEnd4(head *ListNode, n int) *ListNode {
	len := getLength(head)
	pre := &ListNode{0, head}
	cur := pre
	for i := 0; i < len-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return pre.Next
}

func getLength(head *ListNode) (length int) {
	for head != nil {
		head = head.Next
		length++
	}
	return
}
