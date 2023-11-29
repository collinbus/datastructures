package list

type LinkedList interface {
	Append(element int)
	Prepend(element int)
	Get(index int) int
	Insert(element int, index int)
}

type SinglyLinkedList struct {
	Root *SinglyNode
	Size int
}

type SinglyNode struct {
	Value int
	Next  *SinglyNode
}

func (l *SinglyLinkedList) Append(element int) {
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

func (l *SinglyLinkedList) Prepend(element int) {
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

func (l *SinglyLinkedList) Get(index int) int {
	current := l.Root
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value
}

func (l *SinglyLinkedList) Insert(element int, index int) {
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

func (l *SinglyLinkedList) RemoveAt(index int) {
	current := l.Root
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	l.Size -= 1
}

func (l *SinglyLinkedList) Poll() int {
	currentVal := l.Root.Value
	l.Root = l.Root.Next
	l.Size -= 1
	return currentVal
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

type DoublyLinkedList struct {
	Head *DoublyNode
	Tail *DoublyNode
	Size int
}

func (d *DoublyLinkedList) Append(element int) {
	if d.Size == 0 {
		d.Prepend(element)
		d.Tail = d.Head
		return
	}
	current := d.Head
	for i := 1; i < d.Size; i++ {
		current = current.Next
	}
	current.Next = &DoublyNode{
		Value: element,
		Next:  nil,
		Prev:  current,
	}
	d.Tail = current.Next
	d.Size++
}

func (d *DoublyLinkedList) Prepend(element int) {
	current := &DoublyNode{
		Value: element,
		Next:  nil,
		Prev:  nil,
	}
	if d.Head != nil {
		current.Next = d.Head
		current.Next.Prev = current
	} else {
		d.Tail = current
	}
	d.Head = current
	d.Size++
}

func (d *DoublyLinkedList) Get(index int) int {
	current := d.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value
}

func (d *DoublyLinkedList) Insert(element int, index int) {
	if index == 0 {
		d.Prepend(element)
		return
	} else if index == d.Size-1 {
		d.Append(element)
		return
	}
	current := d.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	node := &DoublyNode{
		Value: element,
		Next:  current,
		Prev:  current.Prev,
	}
	node.Next.Prev = node
	node.Prev.Next = node
	d.Size++
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

type DoublyNode struct {
	Value int
	Next  *DoublyNode
	Prev  *DoublyNode
}
