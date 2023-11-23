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
	if l.root == nil {
		l.root = &node{
			value: element,
			next:  nil,
		}
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
