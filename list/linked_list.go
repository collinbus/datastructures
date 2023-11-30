package list

type LinkedList[T any] interface {
	Append(element T)
	Prepend(element T)
	Get(index int) T
	Insert(element T, index int)
	RemoveAt(index int)
	Poll() T
	Pop() T
}

type SinglyLinkedList[T any] struct {
	Root *SinglyNode[T]
	Size int
}

type SinglyNode[T any] struct {
	Value T
	Next  *SinglyNode[T]
}

func (l *SinglyLinkedList[T]) Append(element T) {
	if l.Size == 0 {
		l.Prepend(element)
		return
	} else {
		current := l.Root
		for i := 1; i < l.Size; i++ {
			current = current.Next
		}
		current.Next = &SinglyNode[T]{
			Value: element,
			Next:  nil,
		}
	}
	l.Size++
}

func (l *SinglyLinkedList[T]) Prepend(element T) {
	var next *SinglyNode[T]
	if l.Root == nil {
		next = nil
	} else {
		next = l.Root
	}
	l.Root = &SinglyNode[T]{
		Value: element,
		Next:  next,
	}
	l.Size++
}

func (l *SinglyLinkedList[T]) Get(index int) T {
	current := l.Root
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value
}

func (l *SinglyLinkedList[T]) Insert(element T, index int) {
	if l.Size == 0 {
		l.Prepend(element)
		return
	}
	current := l.Root
	for i := 0; i < index; i++ {
		current = current.Next
	}
	current.Next = &SinglyNode[T]{
		Value: current.Value,
		Next:  current.Next,
	}
	current.Value = element
	l.Size++
}

func (l *SinglyLinkedList[T]) RemoveAt(index int) {
	if index == 0 {
		l.Poll()
		return
	}
	current := l.Root
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	l.Size -= 1
}

func (l *SinglyLinkedList[T]) Poll() T {
	if l.Size == 1 {
		return l.Pop()
	}
	currentVal := l.Root.Value
	l.Root = l.Root.Next
	l.Size -= 1
	return currentVal
}

func (l *SinglyLinkedList[T]) Pop() T {
	if l.Size == 1 {
		current := l.Root.Value
		l.Root = nil
		l.Size -= 1
		return current
	}
	current := l.Root
	for i := 0; i < l.Size-2; i++ {
		current = current.Next
	}
	value := current.Next.Value
	current.Next = nil
	l.Size -= 1
	return value
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

type DoublyLinkedList[T any] struct {
	Head *DoublyNode[T]
	Tail *DoublyNode[T]
	Size int
}

func (d *DoublyLinkedList[T]) Append(element T) {
	if d.Size == 0 {
		d.Prepend(element)
		d.Tail = d.Head
		return
	}
	current := d.Head
	for i := 1; i < d.Size; i++ {
		current = current.Next
	}
	current.Next = &DoublyNode[T]{
		Value: element,
		Next:  nil,
		Prev:  current,
	}
	d.Tail = current.Next
	d.Size++
}

func (d *DoublyLinkedList[T]) Prepend(element T) {
	current := &DoublyNode[T]{
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

func (d *DoublyLinkedList[T]) Get(index int) T {
	current := d.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value
}

func (d *DoublyLinkedList[T]) Insert(element T, index int) {
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
	node := &DoublyNode[T]{
		Value: element,
		Next:  current,
		Prev:  current.Prev,
	}
	node.Next.Prev = node
	node.Prev.Next = node
	d.Size++
}

func (l *DoublyLinkedList[T]) RemoveAt(index int) {
	if index == 0 {
		l.Poll()
		return
	} else if index == (l.Size - 1) {
		l.Pop()
		return
	}
	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	current.Next.Prev = current
	l.Size -= 1
}

func (l *DoublyLinkedList[T]) Poll() T {
	if l.Size == 1 {
		return l.Pop()
	}
	currentVal := l.Head.Value
	l.Head = l.Head.Next
	l.Head.Prev = nil
	l.Size -= 1
	return currentVal
}

func (l *DoublyLinkedList[T]) Pop() T {
	if l.Size == 1 {
		current := l.Head.Value
		l.Head = nil
		l.Tail = nil
		l.Size -= 1
		return current
	}
	current := l.Head
	for i := 0; i < l.Size-2; i++ {
		current = current.Next
	}
	value := current.Next.Value
	current.Next = nil
	l.Tail = current
	l.Size -= 1
	return value
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

type DoublyNode[T any] struct {
	Value T
	Next  *DoublyNode[T]
	Prev  *DoublyNode[T]
}
