package Basic

type Comparable func(c1 interface{}, c2 interface{}) bool

type BinaryTree struct {
	node     interface{}
	left     *BinaryTree
	right    *BinaryTree
	lessFunc Comparable
}

func New(compareFunc Comparable) *BinaryTree {
	return &BinaryTree{
		node:     nil,
		lessFunc: compareFunc,
	}
}

func (tree *BinaryTree) Insert(value interface{}) {
	if tree.node == nil {
		tree.node = value
		tree.right = New(tree.lessFunc)
		tree.left = New(tree.lessFunc)
		return
	}
	if tree.lessFunc(value, tree.node) {
		tree.left.Insert(value)
	} else {
		tree.right.Insert(value)
	}
}

func (tree *BinaryTree) Search(value interface{}) *BinaryTree {
	if tree.node == nil {
		return nil
	}
	if tree.node == value {
		return tree
	}
	if tree.lessFunc(value, tree.node) {
		return tree.left.Search(value)
	} else {
		return tree.right.Search(value)
	}
}

func (tree *BinaryTree) Max() interface{} {
	if tree.node == nil || tree.right.node == nil {
		return tree.node
	}
	return tree.right.Max()
}

func (tree *BinaryTree) Min() interface{} {
	if tree.node == nil || tree.left.node == nil {
		return tree.node
	}
	return tree.left.Min()
}
