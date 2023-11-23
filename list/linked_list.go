package list

type LinkedList struct {
	Root *SinglyNode
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

type SinglyNode struct {
	Value int
	Next  *SinglyNode
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
		current.Next = &SinglyNode{
			Value: element,
			Next:  nil,
		}
	}
	l.Size++
}

func (l *LinkedList) Prepend(element int) {
	var next *SinglyNode
	if l.Root == nil {
		next = nil
	} else if l.Root.Next != nil {
		next = l.Root.Next
	} else {
		next = l.Root
	}
	l.Root = &SinglyNode{
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
	current.Next = &SinglyNode{
		Value: current.Value,
		Next:  current.Next,
	}
	current.Value = element
	l.Size++
}
