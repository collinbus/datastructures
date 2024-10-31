package main

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

type List interface {
	Insert(element, index int)
}

func NewNode(element int) *Node {
	return &Node{Value: element}
}

func NewLinkedList(items []int) *LinkedList {
	list := &LinkedList{}
	for i := 0; i < len(items); i++ {
		list.Insert(items[i], i)
	}
	return list
}

func (l *LinkedList) Insert(element, index int) {
	if l.Head == nil {
		node := NewNode(element)
		l.Head = node
		return
	}
	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	node := NewNode(element)
	if current.Next != nil {
		node.Next = current.Next
	}
	current.Next = node
}
