package list

type LinkedList struct {
	root *node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

type node struct {
	value int
	next  *node
}

func (l *LinkedList) Append(element int) {
	if l.size == 0 {
		l.Prepend(element)
		return
	} else {
		current := l.root
		for i := 1; i < l.size; i++ {
			current = current.next
		}
		current.next = &node{
			value: element,
			next:  nil,
		}
	}
	l.size++
}

func (l *LinkedList) Prepend(element int) {
	var next *node
	if l.root == nil {
		next = nil
	} else if l.root.next != nil {
		next = l.root.next
	} else {
		next = l.root
	}
	l.root = &node{
		value: element,
		next:  next,
	}
	l.size++
}

func (l *LinkedList) Get(index int) int {
	current := l.root
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value
}

func (l *LinkedList) Insert(element int, index int) {
	current := l.root
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.next = &node{
		value: current.value,
		next:  current.next,
	}
	current.value = element
	l.size++
}
