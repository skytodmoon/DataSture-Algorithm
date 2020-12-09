package main

func main() {

}

type ListNode struct {
	Val  string
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next= node.Next.Next

}
