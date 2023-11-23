package list

type LinkedList struct {
	Root *Node
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

type Node struct {
	Value int
	Next  *Node
}

func (l *LinkedList) Append(element int) {
	if l.Size == 0 {
		l.Prepend(element)
		return
	} else {
		current := l.Root
		for i := 1; i < l.Size; i++ {
			current = current.Next
		}
		current.Next = &Node{
			Value: element,
			Next:  nil,
		}
	}
	l.Size++
}

func (l *LinkedList) Prepend(element int) {
	var next *Node
	if l.Root == nil {
		next = nil
	} else if l.Root.Next != nil {
		next = l.Root.Next
	} else {
		next = l.Root
	}
	l.Root = &Node{
		Value: element,
		Next:  next,
	}
	l.Size++
}

func (l *LinkedList) Get(index int) int {
	current := l.Root
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value
}

func (l *LinkedList) Insert(element int, index int) {
	if l.Size == 0 {
		l.Prepend(element)
		return
	}
	current := l.Root
	for i := 0; i < index; i++ {
		current = current.Next
	}
	current.Next = &Node{
		Value: current.Value,
		Next:  current.Next,
	}
	current.Value = element
	l.Size++
}
