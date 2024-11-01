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
	ReomveAt(index int)
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
	node := NewNode(element)
	if l.Head == nil {
		l.Head = node
		return
	}
	if index == 0 {
		node.Next = l.Head
		l.Head = node
		return
	}
	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	if current.Next != nil {
		node.Next = current.Next
	}
	current.Next = node
}

func (l *LinkedList) RemoveAt(index int) {
	current := l.Head
	if index == 0 {
		l.Head = current.Next
	}

	for i := 0; i < index - 1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
}
