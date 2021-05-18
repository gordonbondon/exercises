package bst_iterator

type BSTIterator struct {
	pointer *TreeNode
	stack   []*TreeNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewBSTIterator(root *TreeNode) *BSTIterator {
	stack := make([]*TreeNode, 0)

	start := &TreeNode{
		Val:   -1,
		Right: root,
	}

	iterator := &BSTIterator{
		pointer: start,
		stack:   stack,
	}

	return iterator
}

func (iter *BSTIterator) Next() int {
	if iter.pointer.Right != nil {
		iter.pointer = iter.pointer.Right

		return iter.left()
	}

	return iter.right()
}

func (iter *BSTIterator) HasNext() bool {
	if iter.pointer.Right != nil || len(iter.stack) > 0 {
		return true
	}

	return false
}

func (iter *BSTIterator) left() int {
	for iter.pointer.Left != nil {
		iter.stack = append(iter.stack, iter.pointer)

		iter.pointer = iter.pointer.Left
	}

	return iter.pointer.Val
}

func (iter *BSTIterator) right() int {
	iter.pointer = iter.stack[len(iter.stack)-1]
	iter.stack = iter.stack[:len(iter.stack)-1]

	return iter.pointer.Val
}
