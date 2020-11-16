package main

import (
	"fmt"
)

//Node binary tree
type Node struct {
	Value       int
	Left, Right *Node
}

//Print node's Value
func (node *Node) Print() {
	fmt.Println(node.Value, " ")
}

//SetValue node
func (node *Node) SetValue(v int) {
	if node == nil {
		fmt.Println("Can not setting nil node's value")
		return
	}
	node.Value = v
}

//PreOrder Traverse
func (node *Node) PreOrder() {
	if node == nil {
		return
	}
	node.Print()
	node.Left.PreOrder()
	node.Right.PreOrder()
}

//MidOrder Traverse
func (node *Node) MidOrder() {
	if node == nil {
		return
	}
	node.Left.MidOrder()
	node.Print()
	node.Right.MidOrder()
}

//PostOrder Traverse
func (node *Node) PostOrder() {
	if node == nil {
		return
	}
	node.Left.PostOrder()
	node.Right.PostOrder()
	node.Print()
}

//LevelOrder Traverse
func (node *Node) LevelOrder() {
	if node == nil {
		return
	}
	result := []int{}
	nodes := []*Node{node}
	for len(nodes) > 0 {
		curNode := nodes[0]
		nodes = nodes[1:]
		fmt.Println("Nodes' count:", len(nodes))
		result = append(result, curNode.Value)
		if curNode.Left != nil {
			fmt.Println("in left node")
			nodes = append(nodes, curNode.Left)
			fmt.Println("Nodes' count:", len(nodes))
		}
		if curNode.Right != nil {
			fmt.Println("in right node")
			nodes = append(nodes, curNode.Right)
			fmt.Println("Nodes' count:", len(nodes))
		}
	}

	for _, v := range result {
		fmt.Print(v, " ")
	}

}

//Layers Tree length
func (node *Node) Layers() int {
	if node == nil {
		return 0
	}
	leftLayers := node.Left.Layers()
	rightLayers := node.Right.Layers()
	if leftLayers > rightLayers {
		return leftLayers + 1
	} else {
		return rightLayers + 1
	}
}

//CreateNode init
func CreateNode(v int) *Node {
	return &Node{Value: v}
}

func main() {
	root := &Node{Value: 3}
	root.Left = &Node{}
	root.Left.SetValue(0)
	root.Left.Right = CreateNode(2)
	root.Right = &Node{5, nil, nil}
	root.Right.Left = CreateNode(4)

	fmt.Print("\nPre Order:")
	root.PreOrder()
	fmt.Print("\nMid Order:")
	root.MidOrder()
	fmt.Print("\nPost Order:")
	root.PostOrder()
	fmt.Print("\nLevel Order:")
	root.LevelOrder()
	fmt.Print("\ntree length:", root.Layers())

}
