package main

/**
https://leetcode-cn.com/problems/delete-middle-node-lcci/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
func deleteNode(node *ListNode) {
	*node = *node.Next
}

type ListNode struct {
	Var  int
	Next *ListNode
}

func main() {

}
