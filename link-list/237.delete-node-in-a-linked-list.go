package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Var  int
	Next *ListNode
}

//237. 删除链表中的节点
//https://leetcode-cn.com/problems/delete-node-in-a-linked-list/

//方法一：方法一：node表示将被删除的那个节点，所以将指向node节点的指针域，指向node节点的下一个节点的内存地址，即可删除node节点
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

//方法二：将node节点的下一个节点的值复制给node节点，然后node节点的指针域指向node节点的下下一个节点的内存地址，即可通过删除node节点的下一个节点，达到与直接删除node节点相同的效果

func deleteNode2(node *ListNode) {
	*node = *node.Next
}
