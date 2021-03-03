package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{Val: 4, Next: nil}
	second := ListNode{Val: 5, Next: nil}
	third := ListNode{Val: 1, Next: nil}
	last := ListNode{Val: 9, Next: nil}
	head.Next = &second
	second.Next = &third
	third.Next = &last
	deleteNode(&second)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	if node.Next != nil {
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	}
}
