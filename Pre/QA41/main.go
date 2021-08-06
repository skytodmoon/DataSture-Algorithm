package main

import (
	"fmt"
	"sync"
)

func main() {
	//var tree = &ItemBinarySearchTree{}
	var testArray = [7]*int{}
	testArray[0] = *10
	testArray[1] = 5
	testArray[2] = 15
	testArray[3] = 3
	testArray[4] = 7
	testArray[6] = 18

	fmt.Printf("%+v", testArray)
	//for _, temp := range testArray {
	//	tree.Insert(temp)
	//	fmt.Printf("%+v", tree.root)
	//}
	////fmt.Print(tree.root.Left.Val)
	////fmt.Print(tree.root.Left.Left.Val)
	//
	//fmt.Println(rangeSumBST(tree.root, 7, 15))
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var result int
	if root == nil {
		return 0
	}
	if root.Val >= low && root.Val <= high {
		result += root.Val
		result += rangeSumBST(root.Left, low, high)
		result += rangeSumBST(root.Right, low, high)
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ItemBinarySearchTree struct {
	root *TreeNode
	lock sync.RWMutex
}

// 向二叉搜索树的合适位置插入节点
func (tree *ItemBinarySearchTree) Insert(value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	newNode := &TreeNode{value, nil, nil}
	// 初始化树
	if tree.root == nil {
		tree.root = newNode
	} else {
		// 在树中递归查找正确的位置并插入
		insertNode(tree.root, newNode)
	}
}

func insertNode(node, newNode *TreeNode) {
	fmt.Println(node.Val)
	// 插入到左子树
	if newNode.Val < node.Val {
		if node.Left == nil {
			node.Left = newNode
		} else {
			// 递归查找左边插入
			insertNode(node.Left, newNode)
		}
	} else {
		// 插入到右子树
		if node.Right == nil {
			node.Right = newNode
		} else {
			// 递归查找右边插入
			insertNode(node.Right, newNode)
		}
	}
}
