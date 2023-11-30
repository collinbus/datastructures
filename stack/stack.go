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
	s.elements.Prepend(element)
}

func (s *Stack) Peek() int {
	return s.elements.Get(0)
}

func (s *Stack) Pop() int {
	return s.elements.Poll()
}

func (s *Stack) Size() int {
	return s.elements.Size
}
