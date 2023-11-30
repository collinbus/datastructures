package stack

import "datastructures/list"

type Stack struct {
	elements *list.SinglyLinkedList
}

func NewStack() *Stack {
	return &Stack{
		elements: list.NewSinglyLinkedList(),
	}
}

func (s *Stack) Push(element int) {
	s.elements.Append(element)
}

func (s *Stack) Peek() int {
	return s.elements.Get(s.elements.Size - 1)
}

func (s *Stack) Pop() int {
	return s.elements.Pop()
}

func (s *Stack) Size() int {
	return s.elements.Size
}
