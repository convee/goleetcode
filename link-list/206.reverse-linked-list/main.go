package main

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 *
 * https://leetcode-cn.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (57.68%)
 * Total Accepted:    34.3K
 * Total Submissions: 59.3K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 反转一个单链表。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL
 * 输出: 5->4->3->2->1->NULL
 *
 * 进阶:
 * 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
 *
 */
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }


方法：新建链表，头节点插入法
新建一个头结点，遍历原链表，把每个节点用头结点插入到新建链表中。最后，新建的链表就是反转后的链表。
*/

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}

//思路：用一个 prev 节点保存向前指针，temp 保存向后的临时指针
//在遍历列表时，将当前节点的 next 指针改为指向前一个元素
//由于节点没有引用其上一个节点，因此必须事先存储其前一个元素。
//在更改引用之前，还需要另一个指针来存储下一个节点。不要忘记在最后返回新的头引用！

func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		// 保存当前head.Next节点，防止重新赋值后被覆盖
		// 一轮之后状态：nil<-1 2->3->4
		//              prev   head
		temp := head.Next
		head.Next = prev
		// pre 移动
		prev = head
		// head 移动
		head = temp
	}
	return prev
}
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList3(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

//解题思路
//1->2->3->4
//2->1->3->4
//3->2->1->4
//4->3->2->1
func reverseList4(head *ListNode) *ListNode {
	p := head
	q := head.Next
	for p.Next != nil {
		p.Next = q.Next
		q.Next = head
		head = q
		q = p.Next
	}
	return head
}
func main() {
	head1 := &ListNode{Val: 1}
	head2 := &ListNode{Val: 2}
	head3 := &ListNode{Val: 3}
	head4 := &ListNode{Val: 4}
	head1.Next = head2
	head2.Next = head3
	head3.Next = head4
	reverseList(head1)
}
