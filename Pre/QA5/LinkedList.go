package main

//Node is an element of a list
type Node struct {
	Data interface{}
	Next *Node
}

//LinkedList is a linked list
type LinkedList struct {
	Head   *Node
	Length int
}

//Method Interface method
type Method interface {
	Push(v interface{}) //add
	Pop() *Node         //pop
	IsNull() bool       //check if is null
}

//CreateNode init
func CreateNode(v interface{}) *Node {
	return &Node{v, nil}
}

//CreateList init
func CreateList() *LinkedList {
	return &LinkedList{CreateNode(nil), 0}
}

//Push method
func (list *LinkedList) Push(v interface{}) {
	node := CreateNode(v)
	if list.IsNull() {
		list.Head.Next = node
	} else {
		cur := list.Head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

//Pop method
func (list *LinkedList) Pop() *Node {
	if list.IsNull() {
		return nil
	}
	tempNode := list.Head.Next
	list.Head.Next = list.Head.Next.Next
	return tempNode
}

//IsNull check
func (list *LinkedList) IsNull() bool {
	if list.Head.Next == nil {
		return true
	}
	return false
}
