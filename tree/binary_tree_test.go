package tree

import "testing"

func TestIsBinaryTreeIsNotFull(t *testing.T) {
	/*
				1
			2		3
		4		5		6
	*/
	rootNode := NewNode[int](1)
	tree := NewBinaryTree(rootNode)
	rootNode.left = NewNode[int](2)
	rootNode.right = NewNode[int](3)
	rootNode.left.left = NewNode[int](4)
	rootNode.left.right = NewNode[int](5)
	rootNode.right.left = NewNode[int](6)

	isFull := tree.IsFull()

	if isFull {
		t.Fatal("tree should not be full")
	}
}

func TestIsBinaryTreeIsFull(t *testing.T) {
	/*
				1
			2		3
		4		5
	*/
	rootNode := NewNode[int](1)
	tree := NewBinaryTree(rootNode)
	rootNode.left = NewNode[int](2)
	rootNode.right = NewNode[int](3)
	rootNode.left.left = NewNode[int](4)
	rootNode.left.right = NewNode[int](5)

	isFull := tree.IsFull()

	if !isFull {
		t.Fatal("tree should be full")
	}
}

func TestIsBinaryTreeIsNotPerfect(t *testing.T) {
	/*
				1
			2		3
		4		5		6
	*/
	rootNode := NewNode[int](1)
	tree := NewBinaryTree(rootNode)
	rootNode.left = NewNode[int](2)
	rootNode.right = NewNode[int](3)
	rootNode.left.left = NewNode[int](4)
	rootNode.left.right = NewNode[int](5)
	rootNode.right.left = NewNode[int](6)

	isPerfect := tree.IsPerfect()

	if !isPerfect {
		t.Fatal("tree should not be perfect")
	}
}

func TestIsBinaryTreeIsPerfect(t *testing.T) {
	/*
				1
			2		3
		4	  5  6     7
	*/
	rootNode := NewNode[int](1)
	tree := NewBinaryTree(rootNode)
	rootNode.left = NewNode[int](2)
	rootNode.right = NewNode[int](3)
	rootNode.left.left = NewNode[int](4)
	rootNode.left.right = NewNode[int](5)
	rootNode.right.left = NewNode[int](6)
	rootNode.right.right = NewNode[int](7)

	isPerfect := tree.IsPerfect()

	if !isPerfect {
		t.Fatal("tree should be perfect")
	}
}
