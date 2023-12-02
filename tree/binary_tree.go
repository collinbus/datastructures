package tree

import "math"

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

func (b *BinaryTree[T]) IsPerfect() bool {
	return isTreePerfect[T](b.root, 1, 1)
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

func isTreePerfect[T any](node *Node[T], level int, depth int) bool {
	if node.left != nil && !isTreePerfect(node.left, level+1, depth) {
		return false
	}
	if node.right != nil && !isTreePerfect(node.right, level+1, depth) {
		return false
	}
	depth = int(math.Max(float64(depth), float64(level)))
	return level == depth || (node.left == nil && node.right == nil)
}
