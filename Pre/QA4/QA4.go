package main

import (
	"fmt"
)

//ListNode is a single linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil

	for cur != nil {
		fmt.Println(cur.Val)
		// pre = cur
		// cur = cur.Next
		// cur.Next = pre

		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

func reverseListBy(k int, node *ListNode) *ListNode {
	// spiltNode used to get the k node
	var spiltNode *ListNode = node

	var i = 0
	for i = 0; i < k-1; i++ {
		if spiltNode != nil {
			spiltNode = spiltNode.Next
		} else {
			break
		}
	}

	if spiltNode == nil {
		fmt.Println("nil spilt node:")
		return reverseList(node)
	}

	//next group's head
	nextListHead := spiltNode.Next
	//seprate the group node
	spiltNode.Next = nil

	newHead := reverseList(node)

	nextHeadTmp := reverseListBy(k, nextListHead)

	node.Next = nextHeadTmp

	return newHead

}

func main() {
	head := new(ListNode)
	head.Val = 1
	ln2 := new(ListNode)
	ln2.Val = 2
	ln3 := new(ListNode)
	ln3.Val = 3
	ln4 := new(ListNode)
	ln4.Val = 4
	ln5 := new(ListNode)
	ln5.Val = 5
	ln6 := new(ListNode)
	ln6.Val = 6
	ln7 := new(ListNode)
	ln7.Val = 7
	ln8 := new(ListNode)
	ln8.Val = 8
	ln9 := new(ListNode)
	ln9.Val = 9

	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5
	ln5.Next = ln6
	ln6.Next = ln7
	ln7.Next = ln8
	ln8.Next = ln9

	// pre := reverseList(head)

	// for pre != nil {
	// 	fmt.Println(pre.Val)
	// 	pre = pre.Next
	// }

	pre2 := reverseListBy(3, head)
	for pre2 != nil {
		fmt.Println("index:", pre2.Val)
		pre2 = pre2.Next
	}

}
