package stack

import "datastructures/list"

type Stack interface {
	Push(element int)
	Peek() int
	Pop() int
	Size() int
	IsEmpty() bool
}

type DynamicStack struct {
	elements *list.SinglyLinkedList[int]
}

func NewStack() *DynamicStack {
	return &DynamicStack{
		elements: list.NewSinglyLinkedList[int](),
	}
}

func (s *DynamicStack) Push(element int) {
	s.elements.Append(element)
}

func (s *DynamicStack) Peek() int {
	return s.elements.Get(s.elements.Size - 1)
}

func (s *DynamicStack) Pop() int {
	return s.elements.Pop()
}

func (s *DynamicStack) Size() int {
	return s.elements.Size
}

func (s *DynamicStack) IsEmpty() bool {
	return s.Size() == 0
}
