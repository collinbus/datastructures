package tree

type BinaryTree[T any] struct {
	root *Node[T]
}

func NewBinaryTree[T any](root *Node[T]) *BinaryTree[T] {
	return &BinaryTree[T]{root: root}
}

type Node[T any] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{value: value}
}

func (b *BinaryTree[T]) IsFull() bool {
	return isTreeFull[T](b.root)
}

func isTreeFull[T any](node *Node[T]) bool {
	if node.left != nil && !isTreeFull(node.left) {
		return false
	}
	if node.right != nil && !isTreeFull(node.right) {
		return false
	}
	return (node.left == nil && node.right == nil) ||
		(node.left != nil && node.right != nil)
}
